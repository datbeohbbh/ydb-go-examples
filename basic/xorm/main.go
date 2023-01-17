package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	_ "github.com/ydb-platform/ydb-go-sdk/v3"

	"xorm.io/xorm"
)

var (
	enginePool *EngineWithMode
	showSQL    bool
	dsn        string
	prefix     string
)

func init() {
	required := []string{"ydb"}
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagSet.Usage = func() {
		out := flagSet.Output()
		_, _ = fmt.Fprintf(out, "Usage:\n%s [options]\n", os.Args[0])
		_, _ = fmt.Fprintf(out, "\nOptions:\n")
		flagSet.PrintDefaults()
	}
	flagSet.StringVar(&dsn,
		"ydb", "",
		"YDB connection string",
	)

	flagSet.BoolVar(&showSQL,
		"showSQL", true,
		"Show the generated SQL")

	flagSet.StringVar(&prefix,
		"prefix", "",
		"tables prefix",
	)

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		flagSet.Usage()
		os.Exit(1)
	}
	flagSet.Visit(func(f *flag.Flag) {
		for i, arg := range required {
			if arg == f.Name {
				required = append(required[:i], required[i+1:]...)
			}
		}
	})
	if len(required) > 0 {
		fmt.Printf("\nSome required options not defined: %v\n\n", required)
		flagSet.Usage()
		os.Exit(1)
	}
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	if dsn == "" {
		panic(fmt.Errorf("dsn can not be empty"))
	}

	enginePool = &EngineWithMode{
		engineCached: make(map[string]*xorm.Engine),
		dsn:          dsn,
		ctx:          ctx,
	}

	log.Println("ok: connected to database")

	defer func() {
		_ = enginePool.Close()
	}()

	err := prepareSchema(ctx)
	if err != nil {
		panic(fmt.Errorf("failed on create tables: %v", err))
	}

	err = fillTableWithData(ctx)
	if err != nil {
		panic(fmt.Errorf("failed on fill data into tables: %v", err))
	}

	err = replaceByFetchData(ctx, fmt.Sprintf("`%s`", path.Join(prefix, "test/episodes")))
	if err != nil {
		panic(fmt.Errorf("failed on replace by fetch data: %v", err))
	}

	err = selectDefault(ctx)
	if err != nil {
		panic(fmt.Errorf("failed on select default: %v", err))
	}

	err = selectScan(ctx)
	if err != nil {
		panic(fmt.Errorf("failed on select scan: %v", err))
	}

	err = joinTable(ctx)
	if err != nil {
		panic(fmt.Errorf("failed on join table: %v", err))
	}

	err = updateTable(ctx)
	if err != nil {
		panic(fmt.Errorf("failed on update table: %v", err))
	}

	err = deleteRecords(ctx)
	if err != nil {
		panic(fmt.Errorf("failed on delete records: %v", err))
	}
}
