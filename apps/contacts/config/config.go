package config

import "github.com/flockstore/mannaiah-backend/common/config"

// Config extends the shared GlobalConfig with service-specific settings.
type Config struct {
	config.GlobalConfig
	config.DatabaseConfig
}
