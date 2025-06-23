package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/takumi616/golang-backend-sample/config"
)

func Open(ctx context.Context, cfg *config.Config) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser,
		cfg.DBPassword, cfg.DBName, cfg.DBSslmode,
	)

	// Note: sql.Open does not establish any actual connections
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		slog.ErrorContext(ctx, "failed to open the database", "error", err)
		return nil, err
	}

	// Set connection pool options
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	// Confirm connection is alive
	if err := db.PingContext(ctx); err != nil {
		slog.ErrorContext(ctx, "connection to the database is not alive", "error", err)
		return nil, err
	}

	return db, nil
}
