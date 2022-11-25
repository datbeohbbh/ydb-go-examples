package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/ydb-platform/ydb-go-sdk/v3"
	"xorm.io/xorm"
)

func main() {
	engine, err := xorm.NewEngine("ydb", "grpc://localhost:2136/local")
	if err != nil {
		panic(err)
	}

	log.Println("ok: connected to database")

	defer func() {
		_ = engine.Close()
	}()

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

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
}
