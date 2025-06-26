package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgconn"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PgxClient is a wrapper around pgxpool.Pool implementing DB interface.
type PgxClient struct {
	// Pool is the internal pgx connection pool.
	Pool *pgxpool.Pool
}

// Connect creates a new PostgreSQL connection pool using pgx.
func Connect(ctx context.Context, cfg Config) (*PgxClient, error) {
	pgxCfg, err := pgxpool.ParseConfig(cfg.DSN)
	if err != nil {
		return nil, err
	}

	pgxCfg.MaxConns = cfg.MaxConns
	pgxCfg.MinConns = cfg.MinConns
	pgxCfg.MaxConnLifetime = time.Duration(cfg.ConnMaxLifetime) * time.Second

	pool, err := pgxpool.NewWithConfig(ctx, pgxCfg)
	if err != nil {
		return nil, err
	}

	return &PgxClient{Pool: pool}, nil
}

// Close shuts down the connection pool gracefully.
func (c *PgxClient) Close() {
	c.Pool.Close()
}

// Exec executes a query without returning rows.
func (c *PgxClient) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return c.Pool.Exec(ctx, sql, args...)
}

// Query executes a query that returns multiple rows.
func (c *PgxClient) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return c.Pool.Query(ctx, sql, args...)
}

// QueryRow executes a query that returns a single row.
func (c *PgxClient) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return c.Pool.QueryRow(ctx, sql, args...)
}
