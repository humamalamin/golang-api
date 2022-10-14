package config

import (
	"os"
	conversionHelpers "test-fullstack/helpers/conversion"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type ContextKey string

type Config struct {
	AppEnv                     string `mapstructure:"APPENV"`
	AppTz                      string `mapstructure:"TZ"`
	AppIsDev                   bool
	DatabaseDriver             string `mapstructure:"RESOURCE_DATABASE_DRIVER"`
	DatabaseName               string `mapstructure:"RESOURCE_DATABASE_MASTER_NAME"`
	DatabaseHost               string `mapstructure:"RESOURCE_DATABASE_MASTER_HOST"`
	DatabaseUsername           string `mapstructure:"RESOURCE_DATABASE_MASTER_USERNAME"`
	DatabasePassword           string `mapstructure:"RESOURCE_DATABASE_MASTER_PASSWORD"`
	DatabasePort               string `mapstructure:"RESOURCE_DATABASE_MASTER_PORT"`
	DatabaseSsl                string `mapstructure:"RESOURCE_DATABASE_MASTER_SSL"`
	DatabaseMaxOpenConnections int    `mapstructure:"RESOURCE_DATABASE_MAX_OPEN_CONNECTIONS"`
	DatabaseMaxIdleConnections int    `mapstructure:"RESOURCE_DATABASE_MAX_IDLE_CONNECTIONS"`
	PortHttpServer             string `mapstructure:"SERVER_HTTP_ADDRESS"`
	DatabaseURL                string `mapstructure:"DATABASE_URL"`
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Error().Err(err).Msg("[NewConfig-1] Failed To Read Config")
		return nil, err
	}

	env := os.Getenv("APPENV")
	if env == "" {
		env = "local"
	}

	openConn, _ := conversionHelpers.StrToInt(os.Getenv("RESOURCE_DATABASE_MAX_OPEN_CONNECTIONS"))
	iddleConn, _ := conversionHelpers.StrToInt(os.Getenv("RESOURCE_DATABASE_MAX_IDLE_CONNECTIONS"))

	cfg := &Config{
		AppEnv:                     env,
		AppTz:                      os.Getenv("TZ"),
		AppIsDev:                   false,
		DatabaseDriver:             os.Getenv("RESOURCE_DATABASE_DRIVER"),
		DatabaseName:               os.Getenv("RESOURCE_DATABASE_NAME"),
		DatabaseHost:               os.Getenv("RESOURCE_DATABASE_HOST"),
		DatabaseUsername:           os.Getenv("RESOURCE_DATABASE_USERNAME"),
		DatabasePassword:           os.Getenv("RESOURCE_DATABASE_PASSWORD"),
		DatabasePort:               os.Getenv("RESOURCE_DATABASE_PORT"),
		DatabaseSsl:                os.Getenv("RESOURCE_DATABASE_SSL"),
		DatabaseMaxOpenConnections: openConn,
		DatabaseMaxIdleConnections: iddleConn,
		PortHttpServer:             os.Getenv("SERVER_HTTP_ADDRESS"),
		DatabaseURL:                os.Getenv("DATABASE_URL"),
	}

	cfg.AppIsDev = cfg.AppEnv == "staging" || cfg.AppEnv == "local" || cfg.AppEnv == "dev"

	return cfg, nil
}
