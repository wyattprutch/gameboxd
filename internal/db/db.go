package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wyattprutch/gameboxd/internal/config"
)

// pgxpool gives connection pool to reuse connections
func Connect(cfg *config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}

	// test connection
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("Failed to ping database: %w", err)
	}

	return pool, nil
}
