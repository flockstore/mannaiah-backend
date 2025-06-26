package database

// Config holds the configuration for the PostgreSQL connection.
type Config struct {
	// DSN is the full database connection string.
	DSN string `validate:"required" default:"postgres://user:pass@localhost:5432/app"`

	// MaxConns is the maximum number of open connections in the pool.
	MaxConns int32 `default:"10"`

	// MinConns is the minimum number of idle connections to maintain.
	MinConns int32 `default:"2"`

	// ConnMaxLifetime is the maximum lifetime of a connection in seconds.
	ConnMaxLifetime int `default:"300"`
}
