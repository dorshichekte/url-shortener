package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"

	config "url-shortener/internal/app/config/env"
	url_repository_postgres "url-shortener/internal/app/repository/postgres/url"
)

// ToDo переделать ошибки
func NewConnection(l *zap.Logger, cfg *config.Env) *Postgres {
	db, err := sql.Open("pgx", cfg.DatabaseDSN)
	if err != nil {
		l.Fatal(err.Error())
		panic(err)
	}

	if err = db.Ping(); err != nil {
		l.Fatal(err.Error())
		panic(err)
	}

	err = applyMigrations(cfg.DatabaseDSN)
	if err != nil {
		l.Fatal(err.Error())
		panic(err)
	}

	return &Postgres{Db: db}
}

func applyMigrations(databaseDSN string) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %v", err)
	}

	migrationDirPath := "file://" + filepath.Join(wd, "migrations")

	m, err := migrate.New(migrationDirPath, databaseDSN)
	if err != nil {
		return fmt.Errorf("failed to initialize migration: %v", err)
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		return fmt.Errorf("failed to apply migrations: %v", err)
	}

	return nil
}

func New(db *sql.DB, config *config.Env) Repositories {
	return Repositories{
		Url: url_repository_postgres.New(db, config),
	}
}
