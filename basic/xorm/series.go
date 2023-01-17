package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sort"
	"time"

	"xorm.io/builder"
	"xorm.io/xorm"
)

func createTable(ctx context.Context, engine *xorm.Engine, tableName string, beans interface{}) error {
	session := engine.NewSession().Context(ctx)
	defer func() {
		_ = session.Close()
	}()

	err := session.DropTable(tableName)
	if err != nil {
		log.Printf("warn: drop %s table failed: %v", tableName, err)
	}

	err = session.CreateTable(beans)
	if err != nil {
		return err
	}

	log.Printf("ok: %s table is created", tableName)

	return nil
}

func prepareSchema(ctx context.Context) error {
	engine, err := enginePool.GetSchemeQueryEngine()
	if err != nil {
		return err
	}

	for tableName, ii := range map[string]interface{}{
		(&Series{}).TableName():       &Series{},
		(&Seasons{}).TableName():      &Seasons{},
		(&Episodes{}).TableName():     &Episodes{},
		(&TestEpisodes{}).TableName(): &TestEpisodes{},
	} {
		err := createTable(ctx, engine, tableName, ii)
		if err != nil {
			return err
		}
	}

	return nil
}

func fillTableWithData(ctx context.Context) error {
	engine, err := enginePool.GetDefaultEngine()
	if err != nil {
		return err
	}
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

	session := engine.NewSession().Context(ctx)
	defer func() {
		_ = session.Close()
	}()

	if err = session.Begin(); err != nil {
		return err
	}
	defer func() {
		_ = session.Rollback()
	}()

	// fill series table
	_, err = session.Table(&Series{}).Insert(seriesDataMap)
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
	_, err = session.Table(&TestEpisodes{}).Replace(episodesDataMap)
	if err != nil {
		return err
	}

	if err := session.Commit(); err != nil {
		return err
	}

	log.Println("ok: fill tables with data")

	return nil
}

// FIXME
func explainQuery(ctx context.Context, tableName string) error {
	engine, err := enginePool.GetExplainQueryEngine()
	if err != nil {
		return err
	}
	log.Println("Explain Query")
	// enable `isAutoClose` to close this session after query.
	session := engine.Context(ctx)

	rows, err := session.Table(&Series{}).Cols("series_id", "title", "release_date").Rows(&Series{})
	// rows, err := session.Table("series").Cols("series_id", "title", "release_date").Query()
	defer func() {
		_ = rows.Close()
	}()
	// auto close session after this
	if err != nil {
		return fmt.Errorf("explain query failed: %w", err)
	}

	if rows.Next() {
		var ast, plan string
		err = rows.Scan(&ast, &plan)
		if err != nil {
			return fmt.Errorf("explain query failed: %v", err)
		}
		log.Println(ast, plan)
	}
	// log.Println(rows)

	return nil
}

func selectDefault(ctx context.Context) error {
	/* 	err := explainQuery(ctx, "series")
	   	if err != nil {
	   		return err
	   	} */

	engine, err := enginePool.GetScanQueryEngine()
	if err != nil {
		return err
	}

	log.Println("Select Default")
	session := engine.NewSession().Context(ctx)
	defer func() {
		_ = session.Close()
	}()

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

func selectScan(ctx context.Context) error {
	engine, err := enginePool.GetScanQueryEngine()
	if err != nil {
		return err
	}
	log.Println("Select Scan")
	// Note: `builder` does not apply quote policy of YQL
	var (
		id         []byte
		seriesIDs  [][]byte
		seasonsIDs [][]byte
	)

	session := engine.NewSession().Context(ctx)
	defer func() {
		_ = session.Close()
	}()

	// getting series ID's
	var tmp interface{}
	has, err := session.
		Table((&Series{}).TableName()).
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
		Table((&Seasons{}).TableName()).
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
		Asc("air_date", "title").
		// Limit(3, 3). // Limit(limitN, offsetN)
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

func joinTable(ctx context.Context) error {
	engine, err := enginePool.GetScanQueryEngine()
	if err != nil {
		return err
	}

	session := engine.Context(ctx)

	logResult := func(msg string, res []map[string][]byte) {
		log.Println(msg)
		log.Printf("got %v records", len(res))
		cols := []string{}
		for col := range res[0] {
			cols = append(cols, col)
		}
		sort.Strings(cols)
		for _, m := range res {
			var str string = ""
			for _, col := range cols {
				str += string(m[col]) + "  "
			}
			log.Println(">", str)
		}
	}

	for _, joinOp := range []string{"INNER", "LEFT", "RIGHT", "FULL"} {
		res, rerr := session.
			Table((&Seasons{}).TableName()).
			Alias("sa").
			Join(joinOp, []string{(&Series{}).TableName(), "sr"}, "sa.series_id = sr.series_id").
			Cols("sa.title", "sr.title", "sr.series_id", "sa.season_id").
			Asc("sr.series_id", "sa.season_id").
			Query()

		if rerr != nil {
			return rerr
		}

		logResult(fmt.Sprintf("%s JOIN", joinOp), res)
	}

	res, err := session.
		Table((&Seasons{}).TableName()).
		Alias("sa").
		Join("LEFT SEMI", []string{(&Series{}).TableName(), "sr"}, "sa.series_id = sr.series_id").
		Cols("sa.title", "sa.season_id").
		Asc("sa.season_id").
		Query()

	if err != nil {
		return err
	}

	logResult(fmt.Sprintf("%s JOIN", "LEFT SEMI"), res)

	return nil
}

func updateTable(ctx context.Context) error {
	engine, err := enginePool.GetDefaultEngine()
	if err != nil {
		return err
	}

	log.Println("Update table")
	session := engine.
		NewSession().
		Context(ctx)

	defer func() {
		_ = session.Close()
	}()

	if err = session.Begin(); err != nil {
		return err
	}
	defer func() {
		_ = session.Rollback()
	}()

	_, err = session.
		Table((&Episodes{}).TableName()).
		Update(map[string]interface{}{
			"title": "test",
			"views": uint64(999),
		}, builder.Gte{"air_date": date("2010-12-31")}.And(builder.Like{"title", "%The%"}))

	if err != nil {
		return err
	}

	if err = session.Commit(); err != nil {
		return err
	}

	// read after commit, if not will cause an error
	rows, err := session.
		Table((&Episodes{}).TableName()).
		Cols("title", "air_date", "views").
		Where("views = ?", sql.Named("views", uint64(999))).
		And("title = ?", sql.Named("title", "test")).
		Rows(&Episodes{})

	if err != nil {
		return err
	}

	defer func() {
		_ = rows.Close()
	}()

	log.Println("ok: after episodes table is updated")
	for rows.Next() {
		var (
			title    string
			air_date time.Time
			views    uint64
		)
		// Scan order must be same as Cols order
		// this case: Cols("title", "air_date", "views")
		if err := rows.Scan(&title, &air_date, &views); err != nil {
			return err
		}
		log.Println(">", title, views, air_date.Format("2006-01-02"))
	}

	return nil
}

func deleteRecords(ctx context.Context) error {
	engine, err := enginePool.GetDefaultEngine()
	if err != nil {
		return err
	}
	log.Println("Delete Records")

	_, err = engine.TransactionContext(
		ctx,
		func(ctx context.Context, session *xorm.Session) (_ interface{}, err error) {
			cnt, err := session.
				Table(&Seasons{}).
				Cols("first_aired").
				Count()
			if err != nil {
				return nil, err
			}
			log.Println(">", "before delete:", cnt)

			_, err = session.
				Table(&Seasons{}).
				Where(builder.Between{
					Col:     "first_aired",
					LessVal: sql.Named("from", date("2007-06-01")),
					MoreVal: sql.Named("to", date("2008-06-01")),
				}).
				In("title", []string{"Season 1", "Season 2", "Season 3", "Season 4", "Season 5"}).
				Delete() // can pass the struct and condition will match for struct's field

			if err != nil {
				return nil, err
			}

			log.Println("ok: deleted all records with first_aired from 2007-06-01 to 2008-06-01 in all seasons")

			// no need for commit or rollback, just return nil error as good final
			return nil, nil
		},
	)

	if err != nil {
		return err
	}

	session := engine.Context(ctx)
	cnt, err := session.Table(&Seasons{}).Cols("first_aired").Count()
	if err != nil {
		return err
	}
	log.Println(">", "after delete:", cnt)

	return nil
}

func replaceByFetchData(ctx context.Context, fromTable string) error {
	engine, err := enginePool.GetDefaultEngine()
	if err != nil {
		return err
	}
	// Fetch data from `test/episodes` to `episodes`.
	// After fetch, the number of rows in `episodes` equals the number of rows in `test/episodes`
	log.Println("Replace by fetch data")

	session := engine.NewSession().Context(ctx)
	defer func() {
		_ = session.Close()
	}()

	if err = session.Begin(); err != nil {
		return err
	}
	defer func() {
		_ = session.Rollback()
	}()

	rowsBefore, err := session.
		Table(fromTable).
		Count()

	if err != nil {
		return err
	}
	log.Printf("> number of rows of %s: %v\n", fromTable, rowsBefore)

	// replace by fetch data
	_, err = session.
		Table(&Episodes{}).
		Replace(builder.Select().From(fromTable))

	if err != nil {
		return err
	}

	if err = session.Commit(); err != nil {
		return err
	}

	rowsAfter, err := session.
		Table(&Episodes{}).
		Count()

	if err != nil {
		return err
	}

	log.Println(">", "number of rows after replace:", rowsAfter)

	if rowsAfter != rowsBefore {
		return fmt.Errorf("expected number of [number of rows in %s] = [number of rows in %s]", "episodes", fromTable)
	}

	return nil
}
