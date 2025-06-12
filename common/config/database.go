package config

// DatabaseConfig defines connection settings for a PostgreSQL-compatible database.
type DatabaseConfig struct {
	// DatabaseURL is the full connection string (DSN) to the target database.
	// Example: "postgres://user:password@host:5432/dbname?sslmode=disable"
	DatabaseURL string `mapstructure:"database_url"`

	// MaxPool sets the maximum number of open connections in the connection pool.
	MaxPool int `mapstructure:"db_max_pool"`

	// MinIdle sets the minimum number of idle connections maintained in the pool.
	MinIdle int `mapstructure:"db_min_idle"`

	// MaxConnLifetime defines the maximum amount of time a connection may be reused.
	// Represented in seconds. Use 0 to disable connection lifetime limit.
	MaxConnLifetime int `mapstructure:"db_max_conn_lifetime"`

	// Debug enables or disables SQL debug logging.
	Debug bool `mapstructure:"db_debug"`
}
