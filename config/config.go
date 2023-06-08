package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment               string // develop, staging, production
	LogLevel                  string // DEBUG, INFO ...
	HTTPPort                  string
	PostgresHost              string
	PostgresPort              string
	PostgresDatabase          string
	PostgresUser              string
	PostgresPassword          string
	PostgresConnectionTimeOut int // seconds
	PostgresConnectionTry     int
	BaseUrl                   string
	AdminUsername             string
	AdminPassword             string
}

// Load loads environment vars and inflates Config
func Load() Config {
	dotenvFilePath := cast.ToString(getOrReturnDefault("DOT_ENV_PATH", "config/test.env"))
	err := godotenv.Load(dotenvFilePath)

	if err != nil {
		fmt.Println(".env file not found. Default configuration is being used.")
	}
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "DEBUG"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", "8000"))
	c.BaseUrl = cast.ToString(getOrReturnDefault("BASE_URL", "http://localhost:8000/"))

	// Postgres
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "templatedatabase"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "templateuser"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "templatesecret"))
	c.PostgresConnectionTimeOut = cast.ToInt(getOrReturnDefault("POSTGRES_CONNECTION_TIMEOUT", 5))
	c.PostgresConnectionTry = cast.ToInt(getOrReturnDefault("POSTGRES_CONNECTION_TRY", 10))

	c.AdminUsername = cast.ToString(getOrReturnDefault("ADMIN_USERNAME", "templatename"))
	c.AdminPassword = cast.ToString(getOrReturnDefault("ADMIN_PASSWORD", "templatepass"))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
