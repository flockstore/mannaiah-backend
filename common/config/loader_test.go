package config

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

// testStruct is a mock struct with nested fields and mapstructure tags.
type testStruct struct {
	StringField string       `mapstructure:"string_field"`
	IntField    int          `mapstructure:"int_field"`
	BoolField   bool         `mapstructure:"bool_field"`
	Nested      nestedStruct `mapstructure:"nested"`
	private     string       // should be ignored
	Untagged    string       // no tag, should be ignored
}

type nestedStruct struct {
	NestedString string `mapstructure:"nested_string"`
	NestedInt    int    `mapstructure:"nested_int"`
}

// TestBindEnvs check if default/environment population function works correctly.
func TestBindEnvs(t *testing.T) {

	t.Setenv("STRING_FIELD", "hello")
	t.Setenv("INT_FIELD", "123")
	t.Setenv("BOOL_FIELD", "true")
	t.Setenv("NESTED_NESTED_STRING", "world")
	t.Setenv("NESTED_NESTED_INT", "42")

	v := viper.New()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	bindEnvs(v, "", reflect.TypeOf(testStruct{}))

	var actual testStruct
	err := v.Unmarshal(&actual)

	assert.NoError(t, err)
	assert.Equal(t, "hello", actual.StringField)
	assert.Equal(t, 123, actual.IntField)
	assert.Equal(t, true, actual.BoolField)
	assert.Equal(t, "world", actual.Nested.NestedString)
	assert.Equal(t, 42, actual.Nested.NestedInt)
	assert.Empty(t, actual.Untagged)
}

// TestBindEnvsIgnoresNonStructs ensures bindEnvs exits early on non-struct types.
func TestBindEnvsIgnoresNonStructs(t *testing.T) {
	v := viper.New()

	// Should do nothing, just run safely.
	bindEnvs(v, "", reflect.TypeOf(42))
	bindEnvs(v, "", reflect.TypeOf("hello"))

	ptr := "text"
	bindEnvs(v, "", reflect.TypeOf(&ptr)) // *string

	// no panic, no error => success
	assert.True(t, true)
}

// TestLoadGlobalFromYAML ensures GlobalConfig is loaded correctly from YAML file.
func TestLoadGlobalFromYAML(t *testing.T) {
	cfg, found, err := Load[GlobalConfig]("testdata/global.yaml")
	assert.NoError(t, err)
	assert.True(t, found)

	assert.Equal(t, "contacts", cfg.ServiceName)
	assert.Equal(t, 8081, cfg.Port)
	assert.Equal(t, "debug", cfg.LogLevel)
	assert.Equal(t, EnvDev, cfg.Env)
}

// TestLoadDatabaseFromYAML ensures DatabaseConfig is loaded correctly from YAML file.
func TestLoadDatabaseFromYAML(t *testing.T) {
	cfg, found, err := Load[DatabaseConfig]("testdata/database.yaml")
	assert.NoError(t, err)
	assert.True(t, found)

	assert.Equal(t, "postgres://user:pass@localhost:5432/mannaiah", cfg.DatabaseURL)
	assert.Equal(t, 25, cfg.MaxPool)
	assert.Equal(t, 5, cfg.MinIdle)
	assert.Equal(t, 600, cfg.MaxConnLifetime)
	assert.Equal(t, true, cfg.Debug)
}

// TestLoadGlobalDefaults ensures defaults are applied when YAML omits optional fields.
func TestLoadGlobalDefaults(t *testing.T) {
	cfg, found, err := Load[GlobalConfig]("testdata/global_defaults.yaml")
	assert.NoError(t, err)
	assert.False(t, found)

	assert.Equal(t, "mannaiah-unknown", cfg.ServiceName)
	assert.Equal(t, 8080, cfg.Port)
	assert.Equal(t, "info", cfg.LogLevel)
	assert.Equal(t, EnvDev, cfg.Env)
}

// TestLoadDatabaseDefaults ensures defaults are applied when YAML omits optional fields.
func TestLoadDatabaseDefaults(t *testing.T) {
	cfg, found, err := Load[DatabaseConfig]("testdata/database_defaults.yaml")
	assert.NoError(t, err)
	assert.False(t, found)

	assert.Equal(t, 5, cfg.MinIdle)
	assert.Equal(t, 600, cfg.MaxConnLifetime)
}

// TestLoadGlobalValidationFails ensures GlobalConfig validation fails if required fields are missing.
func TestLoadGlobalValidationFails(t *testing.T) {
	_, _, err := Load[GlobalConfig]("testdata/global_invalid.yaml")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")
}

// TestLoadDatabaseValidationFails ensures DatabaseConfig validation fails if required fields are missing.
func TestLoadDatabaseValidationFails(t *testing.T) {
	_, _, err := Load[DatabaseConfig]("testdata/database_invalid.yaml")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")
}

// TestLoadGlobalFromEnvOnly ensures GlobalConfig loads correctly from environment variables alone.
func TestLoadGlobalFromEnvOnly(t *testing.T) {
	t.Setenv("SERVICE_NAME", "env-service")
	t.Setenv("PORT", "9090")
	t.Setenv("LOG_LEVEL", "info")
	t.Setenv("APP_ENV", "production")

	cfg, found, err := Load[GlobalConfig]("nonexistent.yaml")
	assert.NoError(t, err)
	assert.False(t, found)

	assert.Equal(t, "env-service", cfg.ServiceName)
	assert.Equal(t, 9090, cfg.Port)
	assert.Equal(t, "info", cfg.LogLevel)
	assert.Equal(t, EnvProduction, cfg.Env)
}

// TestLoadDatabaseFromEnvOnly ensures DatabaseConfig loads correctly from environment variables alone.
func TestLoadDatabaseFromEnvOnly(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://env:env@localhost:5432/env")
	t.Setenv("DB_MAX_POOL", "50")
	t.Setenv("DB_MIN_IDLE", "10")
	t.Setenv("DB_MAX_CONN_LIFETIME", "300")
	t.Setenv("DB_DEBUG", "true")

	cfg, found, err := Load[DatabaseConfig]("nonexistent.yaml")
	assert.NoError(t, err)
	assert.False(t, found)

	assert.Equal(t, "postgres://env:env@localhost:5432/env", cfg.DatabaseURL)
	assert.Equal(t, 50, cfg.MaxPool)
	assert.Equal(t, 10, cfg.MinIdle)
	assert.Equal(t, 300, cfg.MaxConnLifetime)
	assert.Equal(t, true, cfg.Debug)
}

// TestLoadFailsOnMalformedYAML checks if Load fails when YAML has invalid types.
func TestLoadFailsOnMalformedYAML(t *testing.T) {
	_, _, err := Load[brokenConfig]("testdata/broken.yaml")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal")
}
