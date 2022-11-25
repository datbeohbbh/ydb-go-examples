package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"xorm.io/builder"
	"xorm.io/xorm"
)

// 25.11.2022:
// TODO: solve problem with transactions

// note: engine.Context() set `isAutoClose` = true in new session.
// to disable `isAutoClose` in new session use: engine.NewSession().Context(ctx) instead (session for transactions must be use this).

func createTable(ctx context.Context, engine *xorm.Engine, tableName string, beans interface{}) error {
	session := engine.NewSession().Context(ydb.WithQueryMode(ctx, ydb.SchemeQueryMode))
	defer session.Close()

	err := session.DropTable(tableName)
	if err != nil {
		log.Printf("warn: drop %s table failed: %v", tableName, err)
	}

	err = session.CreateTable(beans)
	if err != nil {
		log.Printf("error: create %s table failed: %v", tableName, err)
		return err
	}

	log.Printf("ok: %s table is created", tableName)
	return nil
}

func prepareSchema(ctx context.Context, engine *xorm.Engine) error {
	for tableName, ii := range map[string]interface{}{
		(&Series{}).TableName():   &Series{},
		(&Seasons{}).TableName():  &Seasons{},
		(&Episodes{}).TableName(): &Episodes{},
	} {
		err := createTable(ctx, engine, tableName, ii)
		if err != nil {
			return err
		}
	}

	return nil
}

func fillTableWithData(ctx context.Context, engine *xorm.Engine) error {
	// seriesData, seasonsData, episodesData := getData()
	seriesData, seasonsData, episodesData := getData()

	seriesDataMap := []map[string]interface{}{}
	for _, v := range seriesData {
		seriesDataMap = append(seriesDataMap, map[string]interface{}{
			"series_id":    v.SeriesID,
			"title":        v.Title,
			"series_info":  v.SeriesInfo,
			"release_date": v.ReleaseDate,
			"comment":      v.Comment,
		})
	}

	episodesDataMap := []map[string]interface{}{}
	for _, v := range episodesData {
		episodesDataMap = append(episodesDataMap, map[string]interface{}{
			"series_id":  v.SeriesID,
			"season_id":  v.SeasonID,
			"episode_id": v.EpisodeID,
			"title":      v.Title,
			"air_date":   v.AirDate,
		})
	}

	session := engine.NewSession().Context(ydb.WithTxControl(ctx, table.DefaultTxControl())) // default mode
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}
	defer session.Rollback()
	// fill series table
	_, err := session.Table((&Series{}).TableName()).Insert(seriesDataMap)
	if err != nil {
		log.Printf("error: insert data into %s table failed", (&Series{}).TableName())
		return err
	}

	// fill seasons table
	_, err = session.Insert(seasonsData)
	if err != nil {
		log.Printf("error: insert data into %s table failed", (&Seasons{}).TableName())
		return err
	}

	// fill episodes table

	// replace using []map[string]interface{}
	_, err = session.Table((&Episodes{}).TableName()).Replace(episodesDataMap)
	if err != nil {
		return err
	}

	// replace using fetch
	// TODO: this way have not work yet
	/* 	_, err = session.Table((&Episodes{}).TableName()).Replace(builder.Select().From("`dir1/dir2/episodes`"))

	   	if err != nil {
	   		log.Printf("error: replace data into %s table by fetch data from %s table failed", (&Episodes{}).TableName(), "`dir1/dir2/episodes`")
	   		return err
	   	}
	*/
	if err := session.Commit(); err != nil {
		return err
	}
	/*
		log.Println("test after commit")

		rows, err := session.
			Cols("episode_id", "title", "air_date").
			And(builder.Between{
				Col:     "air_date",
				LessVal: sql.Named("from", date("2006-01-01")),
				MoreVal: sql.Named("to", date("2006-12-31")),
			}).
			Rows(&Episodes{})

		defer func() {
			_ = rows.Close()
		}()

		if err != nil {
			return err
		}

		log.Println("> scan select of episodes of `Season 1` of `IT Crowd` between 2006-01-01 and 2006-12-31:")
		for rows.Next() {
			var ep Episodes
			if err = rows.Scan(&ep); err != nil {
				return err
			}
			log.Printf(
				"> [%s] %s (%s)",
				string(ep.EpisodeID), ep.Title, ep.AirDate.Format("2006-01-02"),
			)
		}
	*/
	return nil
}

func TestSelect(ctx context.Context, engine *xorm.Engine) error {
	var series_id interface{}
	session := engine.Context(ydb.WithQueryMode(ctx, ydb.DataQueryMode))
	defer session.Close()

	has, err := session.Table(&Series{}).Cols("series_id").Get(&series_id)
	if err != nil {
		return err
	}

	b, _ := series_id.([]byte)
	log.Println(has, string(b))

	_, err = session.Table(&Series{}).Where("series_id = ?", b).Exist()
	if err != nil {
		return err
	}

	var seasons []Seasons
	err = session.Table(&Seasons{}).Find(&seasons)

	if err != nil {
		return err
	}

	log.Println("Find test")
	for _, v := range seasons {
		log.Println(v.FirstAired, v.LastAired, string(v.SeriesID), string(v.SeasonID), v.Title)
	}

	rows, err := session.Rows(&Seasons{Title: "Season 1"})
	defer func() {
		_ = rows.Close()
	}()

	if err != nil {
		return err
	}

	log.Println("Rows test")
	for rows.Next() {
		var v Seasons
		err = rows.Scan(&v)
		if err != nil {
			return err
		}
		log.Println(v.FirstAired, v.LastAired, string(v.SeriesID), string(v.SeasonID), v.Title)
	}

	return nil
}

func explainQuery(ctx context.Context, engine *xorm.Engine, tableName string) error {
	// enable `isAutoClose` to close this session after query.
	session := engine.Context(ydb.WithQueryMode(ctx, ydb.ExplainQueryMode))

	res, err := session.Table("series").Cols("series_id", "title", "release_date").QueryString()
	// auto close session after this
	if err != nil {
		return fmt.Errorf("explain query failed: %w", err)
	}

	// AST, PLAN
	for i := 0; i < len(res); i++ {
		for k, v := range res[i] {
			log.Println(k, v)
		}
	}
	return nil
}

func selectDefault(ctx context.Context, engine *xorm.Engine) error {
	err := explainQuery(ctx, engine, "series")
	if err != nil {
		return err
	}

	session := engine.NewSession().Context(ydb.WithQueryMode(ctx, ydb.ScanQueryMode))
	defer session.Close()

	rows, err := session.Cols("series_id", "title", "release_date").Rows(&Series{})
	defer func() {
		_ = rows.Close()
	}()

	if err != nil {
		return err
	}

	log.Println("> select of all known series:")
	for rows.Next() {
		var v Series
		if err = rows.Scan(&v); err != nil {
			return err
		}

		log.Printf(
			"> [%s] %s (%s)",
			string(v.SeriesID), v.Title, v.ReleaseDate.Format("2006-01-02"),
		)
	}

	return nil
}

func selectScan(ctx context.Context, engine *xorm.Engine) error {
	// Note: `builder` does not apply quote policy of YQL
	var (
		id         []byte
		seriesIDs  [][]byte
		seasonsIDs [][]byte
	)

	session := engine.NewSession().Context(ydb.WithQueryMode(ctx, ydb.ScanQueryMode))
	defer session.Close()

	// getting series ID's
	var tmp interface{}
	has, err := session.
		Table("series").
		Cols("series_id").
		Where(builder.Like{"title", "%IT Crowd%"}).
		Get(&tmp)

	if err != nil {
		return err
	}

	id, ok := tmp.([]byte)
	if !ok {
		return errors.New("expected []byte")
	}

	if has {
		seriesIDs = append(seriesIDs, id)
	}

	// getting season ID's
	seasonIDsMap, err := session.
		Table("seasons").
		Cols("season_id").
		Where(builder.Like{"title", "%Season 1%"}).
		Query()

	if err != nil {
		return err
	}

	for _, r := range seasonIDsMap {
		id, has := r["season_id"]
		if !has {
			return errors.New("expected `season_id` comlumn")
		}
		seasonsIDs = append(seasonsIDs, id)
	}

	// "getting final query result"
	rows, err := session.
		Cols("episode_id", "title", "air_date").
		In("series_id", seriesIDs).
		In("season_id", seasonsIDs).
		And(builder.Between{
			Col:     "air_date",
			LessVal: sql.Named("from", date("2006-01-01")),
			MoreVal: sql.Named("to", date("2006-12-31")),
		}).
		Rows(&Episodes{})

	defer func() {
		_ = rows.Close()
	}()

	if err != nil {
		return err
	}

	log.Println("> scan select of episodes of `Season 1` of `IT Crowd` between 2006-01-01 and 2006-12-31:")
	for rows.Next() {
		var ep Episodes
		if err = rows.Scan(&ep); err != nil {
			return err
		}
		log.Printf(
			"> [%s] %s (%s)",
			string(ep.EpisodeID), ep.Title, ep.AirDate.Format("2006-01-02"),
		)
	}

	return nil
}
