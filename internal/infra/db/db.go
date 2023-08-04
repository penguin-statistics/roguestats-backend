package db

import (
	"database/sql"

	"exusiai.dev/roguestats-backend/internal/app/appconfig"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func New(conf *appconfig.Config) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(conf.DatabaseURL)))
	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}
