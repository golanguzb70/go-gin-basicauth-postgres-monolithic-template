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
	SMTPEmail                 string
	SMTPEmailPass             string
	SMTPHost                  string
	SMTPPort                  string
}

// Load loads environment vars and inflates Config
func Load() Config {
	dotenvFilePath := cast.ToString(getOrReturnDefault("DOT_ENV_PATH", "config/.env"))
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

	// Email sending
	c.SMTPEmail = cast.ToString(getOrReturnDefault("SMTP_EMAIL", "youremail@gmail.com"))
	c.SMTPEmailPass = cast.ToString(getOrReturnDefault("SMTP_EMAIL_PASS", "YOUR_EMAIL_PASSWORD"))
	c.SMTPHost = cast.ToString(getOrReturnDefault("SMTP_HOST", "smtp host"))
	c.SMTPPort = cast.ToString(getOrReturnDefault("SMTP_PORT", "587"))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
