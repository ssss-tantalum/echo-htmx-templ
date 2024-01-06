package core

import (
	"database/sql"
	"sync"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"

	_ "github.com/go-sql-driver/mysql"
)

var db *bun.DB

func ConnectDB() (*bun.DB, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	var dbOnce sync.Once
	dbOnce.Do(func() {
		sqldb, err := sql.Open("mysql", config.DbDsn)
		if err != nil {
			panic(err)
		}

		db = bun.NewDB(sqldb, mysqldialect.New())
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithEnabled(config.Debug),
			bundebug.FromEnv(""),
		))

	})

	return db, nil
}
