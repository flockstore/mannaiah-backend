package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestConnect_InvalidDSN verifies that an invalid DSN returns an error.
func TestConnect_InvalidDSN(t *testing.T) {
	ctx := context.Background()
	cfg := Config{
		DSN: "invalid-dsn",
	}

	client, err := Connect(ctx, cfg)
	require.Nil(t, client)
	require.Error(t, err)
}

// TestConnect_ValidMockConfig requires a running PostgreSQL instance.
func TestConnect_ValidMockConfig(t *testing.T) {
	t.Skip("Requires a running PostgreSQL database")

	ctx := context.Background()
	cfg := Config{
		DSN:             "postgres://user:pass@localhost:5432/testdb",
		MaxConns:        5,
		MinConns:        1,
		ConnMaxLifetime: 60,
	}

	client, err := Connect(ctx, cfg)
	require.NoError(t, err)
	require.NotNil(t, client)

	client.Close()
}
