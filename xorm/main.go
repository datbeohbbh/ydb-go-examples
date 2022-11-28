package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"xorm.io/xorm"
	"xorm.io/xorm/core"
)

func connect(ctx context.Context) (*xorm.Engine, error) {
	nativeDriver, err := ydb.Open(ctx,
		os.Getenv("YDB_DSN"),
		ydb.WithCertificatesFromFile(os.Getenv("YDB_CERT")),
	)

	if err != nil {
		return nil, err
	}

	connector, err := ydb.Connector(nativeDriver)
	if err != nil {
		return nil, err
	}

	db := sql.OpenDB(connector)

	xdb := core.FromDB(db)
	engine, err := xorm.NewEngineWithDB("ydb", os.Getenv("YDB_DSN"), xdb)
	if err != nil {
		return nil, err
	}
	return engine, nil
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	engine, err := connect(ctx)
	if err != nil {
		panic(err)
	}
	// uncomment to see the generated YQL
	// engine.ShowSQL(true)

	log.Println("ok: connected to database")

	defer func() {
		_ = engine.Close()
	}()

	err = prepareSchema(ctx, engine)
	if err != nil {
		panic(fmt.Errorf("failed on create tables: %v", err))
	}

	err = fillTableWithData(ctx, engine)
	if err != nil {
		panic(fmt.Errorf("failed on fill data into tables: %v", err))
	}

	err = selectDefault(ctx, engine)
	if err != nil {
		panic(fmt.Errorf("failed on select default: %v", err))
	}

	err = selectScan(ctx, engine)
	if err != nil {
		panic(fmt.Errorf("failed on select scan: %v", err))
	}

	err = joinTable(ctx, engine)
	if err != nil {
		panic(fmt.Errorf("failed on join table: %v", err))
	}

	err = updateTable(ctx, engine)
	if err != nil {
		panic(fmt.Errorf("failed on update table: %v", err))
	}

	err = deleteRecords(ctx, engine)
	if err != nil {
		panic(fmt.Errorf("failed on delete records: %v", err))
	}

	err = replaceByFetchData(ctx, engine, "`test/episodes`")
	if err != nil {
		panic(fmt.Errorf("failed on replace by fetch data: %v", err))
	}
}
