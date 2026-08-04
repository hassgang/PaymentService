package db

import (
	"database/sql"
	"fmt"

	"payment_api/pkg/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type DB interface {
	QueryRow(query string, args ...any) *sql.Row
	QueryRowx(query string, args ...any) *sqlx.Row
	Select(dest any, query string, args ...any) error
	Exec(query string, args ...any) (sql.Result, error)
	Ping() error
	Close() error
}

type Database struct {
	*sqlx.DB
}

func New(cfg config.DatabaseConfig) (*Database, error) {
	sqlxDB, err := sqlx.Connect("postgres", cfg.ConnectionString)
	if err != nil {
		return nil, err
	}

	err = sqlxDB.Ping()
	if err != nil {
		return nil, err
	}

	err = goose.SetDialect("postgres")
	if err != nil {
		return nil, err
	}

	err = goose.Up(sqlxDB.DB, cfg.MigrationsPath)
	if err != nil {
		return nil, fmt.Errorf("migrations failed: %w", err)
	}

	return &Database{DB: sqlxDB}, nil
}
