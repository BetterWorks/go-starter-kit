package config

import (
	"log"

	"github.com/spf13/viper"
)

// Configuration defines app configuration on startup
type Configuration struct {
	HttpAPI  HttpAPI  `validate:"required"`
	Logger   Logger   `validate:"required"`
	Postgres Postgres `validate:"required"`
	Services Services `validate:"required"`
}

type HttpAPI struct {
	BaseURL   string
	Namespace string
	Port      uint `validate:"max=65535"`
	Version   uint
}

type Logger struct {
	App LoggerConfig `validate:"required"`
}

type LoggerConfig struct {
	Enabled bool   `validate:"oneof=false true"`
	Label   string `validate:"required"`
	Level   string `validate:"oneof=debug info warn error fatal"`
}

type Postgres struct {
	Database string `validate:"required"`
	Host     string `validate:"required"`
	Password string `validate:"required"`
	Port     uint   `validate:"max=65535"`
	User     string `validate:"required"`
}

type Services struct {
	Example ServiceConfig
}

type ServiceConfig struct {
	BaseURL string
	Timeout uint
}

// LoadConfiguration loads config parameters on startup
func LoadConfiguration() (*Configuration, error) {
	var config Configuration

	viper.SetConfigName("config")

	viper.AddConfigPath("/app/config")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	viper.AllowEmptyEnv(true)

	// http api
	if err := viper.BindEnv("httpapi.baseURL", "HTTPAPI_BASEURL"); err != nil {
		log.Fatalf("error binding env var `HTTPAPI_BASEURL`: %v", err)
	}
	if err := viper.BindEnv("httpapi.port", "HTTPAPI_PORT"); err != nil {
		log.Fatalf("error binding env var `HTTPAPI_PORT`: %v", err)
	}

	// logger
	if err := viper.BindEnv("logger.app.enabled", "LOGGER_APP_ENABLED"); err != nil {
		log.Fatalf("error binding env var `LOGGER_APP_ENABLED`: %v", err)
	}
	if err := viper.BindEnv("logger.app.level", "LOGGER_APP_LEVEL"); err != nil {
		log.Fatalf("error binding env var `LOGGER_APP_LEVEL`: %v", err)
	}

	// postgres
	if err := viper.BindEnv("postgres.database", "POSTGRES_DB"); err != nil {
		log.Fatalf("error binding env var `POSTGRES_DB`: %v", err)
	}
	if err := viper.BindEnv("postgres.host", "POSTGRES_HOST"); err != nil {
		log.Fatalf("error binding env var `POSTGRES_HOST`: %v", err)
	}
	if err := viper.BindEnv("postgres.password", "POSTGRES_PASSWORD"); err != nil {
		log.Fatalf("error binding env var `POSTGRES_PASSWORD`: %v", err)
	}
	if err := viper.BindEnv("postgres.port", "POSTGRES_PORT"); err != nil {
		log.Fatalf("error binding env var `POSTGRES_PORT`: %v", err)
	}
	if err := viper.BindEnv("postgres.user", "POSTGRES_USER"); err != nil {
		log.Fatalf("error binding env var `POSTGRES_USER`: %v", err)
	}

	// service example
	if err := viper.BindEnv("services.example.baseURL", "SVC_EXAMPLE_BASEURL"); err != nil {
		log.Fatalf("error binding env var `SVC_EXAMPLE_BASEURL`: %v", err)
	}
	if err := viper.BindEnv("services.example.timeout", "SVC_EXAMPLE_TIMEOUT"); err != nil {
		log.Fatalf("error binding env var `SVC_EXAMPLE_TIMEOUT`: %v", err)
	}

	// read and unmarshal config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %s", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("error unmarshalling configuration: %v", err)
	}

	return &config, nil
}
