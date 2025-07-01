package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
	"reflect"
	"strings"
)

// validator instance for struct-level validation
var validate = validator.New()

// Default sets default values defined via struct tags.
func Default(cfg any) {
	defaults.SetDefaults(cfg)
}

// bindEnvs binds ENV variables recursively to all fields with a mapstructure tag.
func bindEnvs(v *viper.Viper, prefix string, t reflect.Type) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Skip unexported
		if field.PkgPath != "" {
			continue
		}

		tag := field.Tag.Get("mapstructure")
		if tag == "" || tag == "-" {
			continue
		}

		fullKey := tag
		if prefix != "" {
			fullKey = prefix + "." + tag
		}

		// Recursively bind nested structs
		if field.Type.Kind() == reflect.Struct {
			bindEnvs(v, fullKey, field.Type)
		} else {
			envKey := strings.ToUpper(strings.ReplaceAll(fullKey, ".", "_"))
			_ = v.BindEnv(fullKey, envKey)
		}
	}
}

// Load reads configuration from YAML file and environment variables.
// It unmarshalls the config into type T, applies environment overrides,
// sets defaults, and validates it using validator.v10.
func Load[T any](configPath string) (*T, bool, error) {
	v := viper.New()

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")

	var raw T
	bindEnvs(v, "", reflect.TypeOf(raw))

	fileFound := false
	if err := v.ReadInConfig(); err == nil {
		fileFound = true
	}

	if err := v.Unmarshal(&raw); err != nil {
		return nil, fileFound, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	Default(&raw)

	if err := validate.Struct(&raw); err != nil {
		return nil, fileFound, fmt.Errorf("validation failed: %w", err)
	}

	return &raw, fileFound, nil
}
