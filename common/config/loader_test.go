package config_test

import (
	"testing"

	"github.com/flockstore/mannaiah-backend/common/config"
	"github.com/stretchr/testify/assert"
)

// TestLoadGlobalFromYAML ensures GlobalConfig is loaded correctly from testdata/global.yaml.
func TestLoadGlobalFromYAML(t *testing.T) {
	cfg, found, err := config.Load[config.GlobalConfig]("testdata/global.yaml")
	assert.NoError(t, err)
	assert.True(t, found)

	assert.Equal(t, "contacts", cfg.ServiceName)
	assert.Equal(t, 8081, cfg.Port)
	assert.Equal(t, "debug", cfg.LogLevel)
	assert.Equal(t, config.EnvDev, cfg.Env)
}

// TestLoadDatabaseFromYAML ensures DatabaseConfig is loaded correctly from testdata/database.yaml.
func TestLoadDatabaseFromYAML(t *testing.T) {
	cfg, found, err := config.Load[config.DatabaseConfig]("testdata/database.yaml")
	assert.NoError(t, err)
	assert.True(t, found)

	assert.Equal(t, "postgres://user:pass@localhost:5432/mannaiah", cfg.DatabaseURL)
	assert.Equal(t, 25, cfg.MaxPool)
	assert.Equal(t, 5, cfg.MinIdle)
	assert.Equal(t, 600, cfg.MaxConnLifetime)
	assert.Equal(t, true, cfg.Debug)
}
