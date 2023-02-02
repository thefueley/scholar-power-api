package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	*sql.DB
}

func NewDatabase() (*Database, error) {
	dbConn, err := sql.Open("sqlite3", "internal/db/scholarpower.db")
	if err != nil {
		return &Database{}, fmt.Errorf("could not connect to db: %w", err)
	}

	if err := dbConn.PingContext(context.Background()); err != nil {
		return &Database{}, fmt.Errorf("could not ping db: %w", err)
	}

	return &Database{dbConn}, nil
}

func (d *Database) Ping(ctx context.Context) error {
	return d.DB.PingContext(ctx)
}
