package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	xormLog "xorm.io/xorm/log"

	_ "github.com/ydb-platform/ydb-go-sdk/v3"

	"xorm.io/xorm"
)

type QueryMode int

type EngineWithMode struct {
	engineCached map[string]*xorm.Engine
	dsn          string
	ctx          context.Context
	mu           sync.Mutex
}

const (
	_ = iota
	DataQueryMode
	ExplainQueryMode
	ScanQueryMode
	SchemeQueryMode
	ScriptingQueryMode

	DefaultQueryMode = DataQueryMode
)

var (
	typeToString = map[QueryMode]string{
		DataQueryMode:      "data",
		ScanQueryMode:      "scan",
		ExplainQueryMode:   "explain",
		SchemeQueryMode:    "scheme",
		ScriptingQueryMode: "scripting",
	}
)

func CreateEngine(dsn string) (*xorm.Engine, error) {
	return xorm.NewEngine("ydb", dsn)
}

func (em *EngineWithMode) getEngine(queryMode QueryMode) (*xorm.Engine, error) {
	em.mu.Lock()
	defer em.mu.Unlock()
	mode := typeToString[queryMode]

	if _, has := em.engineCached[mode]; has {
		return em.engineCached[mode], nil
	}

	engine, err := CreateEngine(fmt.Sprintf("%s?query_mode=%s", em.dsn, mode))
	if err != nil {
		return nil, err
	}

	engine.ShowSQL(showSQL)
	engine.SetLogLevel(xormLog.LOG_DEBUG)

	loc, _ := time.LoadLocation("Europe/Moscow")
	engine.SetTZLocation(loc)
	engine.SetTZDatabase(loc)

	engine.SetDefaultContext(em.ctx)

	engine.SetMaxOpenConns(50)
	engine.SetMaxIdleConns(50)
	engine.DB().SetConnMaxIdleTime(time.Second)

	em.engineCached[mode] = engine
	return em.engineCached[mode], nil
}

func (em *EngineWithMode) Close() error {
	em.mu.Lock()
	defer em.mu.Unlock()
	var retErr error = nil
	for mode, engine := range em.engineCached {
		log.Println("Close", mode, "engine")
		if err := engine.Close(); err != nil {
			retErr = err
			break
		}
		delete(em.engineCached, mode)
	}
	return retErr
}

func (em *EngineWithMode) GetDefaultEngine() (*xorm.Engine, error) {
	return em.getEngine(DefaultQueryMode)
}

func (em *EngineWithMode) GetDataQueryEngine() (*xorm.Engine, error) {
	return em.getEngine(DataQueryMode)
}

func (em *EngineWithMode) GetScanQueryEngine() (*xorm.Engine, error) {
	return em.getEngine(ScanQueryMode)
}

func (em *EngineWithMode) GetExplainQueryEngine() (*xorm.Engine, error) {
	return em.getEngine(ExplainQueryMode)
}

func (em *EngineWithMode) GetSchemeQueryEngine() (*xorm.Engine, error) {
	return em.getEngine(SchemeQueryMode)
}

func (em *EngineWithMode) GetScriptQueryEngine() (*xorm.Engine, error) {
	return em.getEngine(ScriptingQueryMode)
}
