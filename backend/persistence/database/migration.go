package database

import (
	"database/sql"
	"errors"
	"github.com/pressly/goose"
)

var (
	ErrMigrationFile = errors.New("please provide a valid name for the migration file")
)

const (
	// Migration directory
	migrationDirectory string = "persistence/database/migrations"

	// Migrations table
	migrationTable string = "daps_db_version"
)

func init() {
	goose.SetTableName(migrationTable)
}

func Migrate(db *sql.DB) error {
	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	if err := goose.Run("up", db, migrationDirectory, "sql"); err != nil {
		return err
	}

	return nil
}

func Rollback(db *sql.DB) error {
	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	if err := goose.Run("down", db, migrationDirectory); err != nil {
		return err
	}
	return nil
}

func MakeMigrations(db *sql.DB, filename string) error {
	if filename == "" {
		return ErrMigrationFile
	}

	if err := goose.Run("create", db, migrationDirectory, filename, "sql"); err != nil {
		return err
	}

	return nil
}
