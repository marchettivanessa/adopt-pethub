package config

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Env       string
	AppPort   string
	Database  DatabaseConfig
	Log       LogConfig
	JWTSecret string
}

type DatabaseConfig struct {
	Name           string
	Host           string
	Password       string
	Username       string
	MigrationPath  string
	Port           int
	ConnectTimeout int
	Schema         string
}

type LogConfig struct {
	AppName            string
	Level              string
	EnableSyslog       bool
	EnableReportCaller bool
}

func MustParseConfig() Config {
	err := godotenv.Load(fmt.Sprintf("%s/.env", os.Getenv("PWD")))
	if err != nil {
		log.WithError(err).Warn("failed to load .env file")
	}
	return Config{
		Env:     MustGetEnv("ENV"),
		AppPort: MustGetEnv("APP_PORT"),
		Database: DatabaseConfig{
			Name:           MustGetEnv("DATABASE_NAME"),
			Host:           MustGetEnv("DATABASE_HOST"),
			Password:       MustGetEnv("DATABASE_PASSWORD"),
			Username:       MustGetEnv("DATABASE_USERNAME"),
			MigrationPath:  MustGetEnv("DATABASE_MIGRATION_PATH"),
			Port:           MustParseInt("DATABASE_PORT"),
			ConnectTimeout: MustParseInt("DATABASE_CONNECTION_TIMEOUT"),
			Schema:         MustGetEnv("DATABASE_SCHEMA"),
		},
		Log: LogConfig{
			AppName:            MustGetEnv("APP_NAME"),
			Level:              MustGetEnv("LOG_LEVEL"),
			EnableSyslog:       MustParseBool(MustGetEnv("LOG_ENABLE_SYSLOG")),
			EnableReportCaller: MustParseBool(MustGetEnv("LOG_ENABLE_REPORT_CALLER")),
		},
		JWTSecret: MustGetEnv("JWT_SECRET"),
	}
}

func SetupCors(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:2000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderContentType},
	}))
}

func MustGetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.WithField("key", key).Fatal("failed to find environment variable")
	}
	return v
}

func MustParseBool(v string) bool {
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.WithError(err).WithField("value", v).Fatal("failed converting value to bool")
	}
	return b
}

func MustParseInt(key string) int {
	v, err := strconv.Atoi(MustGetEnv(key))
	if err != nil {
		log.WithError(err).Fatal("no valid value assigned to env variable", key)
	}
	return v
}
