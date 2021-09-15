package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"text/template"
	"time"

	environ "github.com/ydb-platform/ydb-go-sdk-auth-environ"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/connect"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
)

func render(t *template.Template, data interface{}) string {
	var buf bytes.Buffer
	err := t.Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

type templateConfig struct {
	TablePathPrefix string
}

type result struct {
	code int
	err  string
}

type service struct {
	db       *connect.Connection
	database string
	client   *http.Client
}

func NewService(ctx context.Context, connectParams connect.ConnectParams, connectTimeout time.Duration) (h *service, err error) {
	connectCtx, cancel := context.WithTimeout(ctx, connectTimeout)
	defer cancel()
	db, err := connect.New(connectCtx, connectParams, environ.WithEnvironCredentials(ctx))
	if err != nil {
		return nil, fmt.Errorf("connect error: %w", err)
	}
	h = &service{
		db:       db,
		database: connectParams.Database(),
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
			Timeout: time.Second * 10,
		},
	}
	err = h.createTable(ctx)
	if err != nil {
		defer func() { _ = db.Close() }()
		return nil, fmt.Errorf("error on create table: %w", err)
	}
	return h, nil
}

func (s *service) Close() {
	defer func() { _ = s.db.Close() }()
}

func (s *service) createTable(ctx context.Context) (err error) {
	query := render(
		template.Must(template.New("").Parse(`
			CREATE TABLE healthchecks (
				url         Utf8,
				code        Int32,
				ts          DateTime,
				error       Utf8,

				PRIMARY KEY (url, ts)
			);
		`)),
		templateConfig{
			TablePathPrefix: s.database,
		},
	)
	return table.Retry(ctx, s.db.Table().Pool(),
		table.OperationFunc(func(ctx context.Context, s *table.Session) error {
			err := s.ExecuteSchemeQuery(ctx, query)
			return err
		}),
	)
}

func (s *service) ping(path string) result {
	uri, err := url.Parse(path)
	if err != nil {
		return result{
			code: -1,
			err:  err.Error(),
		}
	}
	if uri.Scheme == "" {
		uri.Scheme = "http"
	}
	request, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return result{
			code: -1,
			err:  err.Error(),
		}
	}
	response, err := s.client.Do(request)
	if err != nil {
		return result{
			code: -1,
			err:  err.Error(),
		}
	}
	return result{
		code: response.StatusCode,
	}
}

func (s *service) check(ctx context.Context, urls []string) (err error) {
	if len(urls) == 0 {
		return nil
	}
	codes := &sync.Map{}
	wg := &sync.WaitGroup{}
	wg.Add(len(urls))
	for _, u := range urls {
		go func(u string) {
			defer wg.Done()
			codes.Store(u, s.ping(u))
		}(u)
	}
	wg.Wait()

	return s.saveCodes(ctx, codes)
}

func (s *service) saveCodes(ctx context.Context, codes *sync.Map) (err error) {
	query := fmt.Sprintf(`
        PRAGMA TablePathPrefix("%s");

		DECLARE $url AS Utf8;
        DECLARE $code AS Int32;
        DECLARE $ts AS DateTime;
        DECLARE $error AS Utf8;

        UPSERT INTO healthchecks ( url, code, ts, error )
        VALUES ($url, $code, $ts, $error);`,
		s.database,
	)
	errs := make(chan error)
	go func() {
		codes.Range(
			func(url, code interface{}) bool {
				res := code.(result)
				writeTx := table.TxControl(table.BeginTx(table.WithSerializableReadWrite()), table.CommitTx())
				err = table.Retry(
					ctx,
					s.db.Table().Pool(),
					table.OperationFunc(
						func(ctx context.Context, s *table.Session) (err error) {
							_, _, err = s.Execute(
								ctx,
								writeTx,
								query,
								table.NewQueryParameters(
									table.ValueParam("$url", ydb.UTF8Value(url.(string))),
									table.ValueParam("$code", ydb.Int32Value(int32(res.code))),
									table.ValueParam("$ts", ydb.DatetimeValueFromTime(time.Now())),
									table.ValueParam("$error", ydb.UTF8Value(res.err)),
								),
							)
							return err
						},
					),
				)
				if err != nil {
					fmt.Println("error on save code", url, res, err)
					errs <- err
				}
				return true
			},
		)
		close(errs)
	}()
	ee := make([]string, 0)
	for err := range errs {
		ee = append(ee, err.Error())
	}
	if len(ee) == 0 {
		return nil
	}
	return fmt.Errorf("errors: [%s]", strings.Join(ee, ","))
}
