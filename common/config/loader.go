package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Load reads configuration from YAML and environment variables.
// It returns the config, a boolean indicating if a config file was found, and an error.
//
// Example:
//
//	cfg, found, err := Load[GlobalConfig]("config.yaml")
func Load[T any](configPath string) (*T, bool, error) {
	v := viper.New()

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")

	fileFound := false
	if err := v.ReadInConfig(); err == nil {
		fileFound = true
	}

	var cfg T

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fileFound, fmt.Errorf("failed to unmarshal structure before binding: %w", err)
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fileFound, fmt.Errorf("failed to unmarshal final config: %w", err)
	}

	return &cfg, fileFound, nil
}
