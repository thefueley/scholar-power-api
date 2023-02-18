package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (d *Database) Migrate() error {
	fmt.Println("Migrating database")

	db, err := sql.Open("sqlite3", "./internal/db/data/scholarpower.db")
	if err != nil {
		fmt.Println("error opening database")
		return err
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		fmt.Println("error creating sqlite3 driver")
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/migrations",
		"sqlite3",
		driver,
	)

	if err != nil {
		fmt.Println("error creating migrate instance")
		return err
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("error migrating database: %w", err)
		}
	}

	fmt.Println("Database migrated successfully")
	return nil
}
