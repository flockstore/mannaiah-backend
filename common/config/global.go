package config

// AppEnv defines the runtime environment in which the application is running.
type AppEnv string

const (
	// EnvLocal represents local development environment.
	EnvLocal AppEnv = "local"

	// EnvDev represents a development (non-prod) deployment.
	EnvDev AppEnv = "dev"

	// EnvStaging represents a staging/test deployment.
	EnvStaging AppEnv = "staging"

	// EnvProduction represents a production deployment.
	EnvProduction AppEnv = "production"
)

// GlobalConfig holds runtime settings common to all Mannaiah microservices.
type GlobalConfig struct {
	// ServiceName is the unique identifier of the running microservice.
	ServiceName string `mapstructure:"service_name"`

	// Port is the port number the service will listen on.
	Port int `mapstructure:"port"`

	// LogLevel defines the verbosity of log output.
	LogLevel string `mapstructure:"log_level"`

	// Env specifies the environment mode used to control behavior like logging, metrics, etc.
	Env AppEnv `mapstructure:"app_env"`
}
