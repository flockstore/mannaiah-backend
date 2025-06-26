package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// DB defines a contract for PostgreSQL operations.
// It can be implemented by real or mock database clients.
type DB interface {
	// Exec executes a query without returning rows.
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)

	// Query executes a query and returns multiple rows.
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)

	// QueryRow executes a query and returns a single row.
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}
