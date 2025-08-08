// Package postgres инициализирует соединение с базой данных PostgreSQL
package postgres

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"

	config "url-shortener/internal/app/config/env"
	url_repository_postgres "url-shortener/internal/app/repository/postgres/url"
	customerror "url-shortener/internal/pkg/error"
)

// NewConnection создает и проверяет соединение с PostgreSQL.
func NewConnection(l *zap.Logger, cfg *config.Env) *sql.DB {
	db, err := sql.Open("pgx", cfg.DatabaseDSN)
	if err != nil {
		l.Error(err.Error())
	}

	if err = db.Ping(); err != nil {
		l.Error(err.Error())
		return nil
	}

	err = applyMigrations(cfg.DatabaseDSN)
	if err != nil {
		l.Fatal(err.Error())
	}

	return db
}

func applyMigrations(databaseDSN string) error {
	wd, err := os.Getwd()
	if err != nil {
		return customerror.NewWithData(errMessageFailedGetCurrentDirectory, err)
	}

	migrationDirPath := "file://" + filepath.Join(wd, "migrations")

	m, err := migrate.New(migrationDirPath, databaseDSN)
	if err != nil {
		return customerror.NewWithData(errMessageFailedInitializeMigrations, err)
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		return customerror.NewWithData(errMessageFailedApplyMigrations, err)
	}

	return nil
}

// New создает структуру репозиториев, инициализируя их подключением к базе.
func New(db *sql.DB, config *config.Env) Repositories {
	return Repositories{
		URL: url_repository_postgres.New(db, config),
	}
}
