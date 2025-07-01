package database

import (
	"context"
	"github.com/flockstore/mannaiah-backend/common/config"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestConnect_InvalidDSN verifies that an invalid DSN returns an error.
func TestConnect_InvalidDSN(t *testing.T) {
	ctx := context.Background()
	cfg := config.DatabaseConfig{
		DatabaseURL: "invalid-dsn",
	}

	client, err := Connect(ctx, cfg)
	require.Nil(t, client)
	require.Error(t, err)
}

// TestConnect_ValidMockConfig requires a running PostgreSQL instance.
func TestConnect_ValidMockConfig(t *testing.T) {
	t.Skip("Requires a running PostgreSQL database")

	ctx := context.Background()
	cfg := config.DatabaseConfig{
		DatabaseURL:     "postgres://user:pass@localhost:5432/testdb",
		MaxPool:         5,
		MinIdle:         1,
		MaxConnLifetime: 60,
	}

	client, err := Connect(ctx, cfg)
	require.NoError(t, err)
	require.NotNil(t, client)

	client.Close()
}
