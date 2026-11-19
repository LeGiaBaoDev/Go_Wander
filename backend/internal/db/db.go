// Package db creates the PostgreSQL connection pool used by the API.
package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Pinger is the subset of the database pool used by readiness checks.
type Pinger interface {
	Ping(ctx context.Context) error
}

// Open creates a connection pool without connecting eagerly.
// An empty URL returns (nil, nil) so the API can still start and report
// itself not ready until the database is configured.
func Open(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	if databaseURL == "" {
		return nil, nil
	}

	cfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("parse DATABASE_URL: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("create database pool: %w", err)
	}
	return pool, nil
}
