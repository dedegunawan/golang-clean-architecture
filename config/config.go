package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName  string
	AppEnv   string
	HTTPAddr string

	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
	DBParams string

	LogLevel string
}

func LoadDotEnv() error {
	return godotenv.Load()
}

func FromEnv() Config {
	return Config{
		AppName:  get("APP_NAME", "golang-clean-architecture"),
		AppEnv:   get("APP_ENV", "development"),
		HTTPAddr: get("HTTP_ADDR", ":8080"),
		DBHost:   get("DB_HOST", "127.0.0.1"),
		DBPort:   get("DB_PORT", "3306"),
		DBUser:   get("DB_USER", "root"),
		DBPass:   get("DB_PASS", ""),
		DBName:   get("DB_NAME", "yourapp"),
		DBParams: get("DB_PARAMS", "charset=utf8mb4&parseTime=True&loc=Local"),
		LogLevel: get("LOG_LEVEL", "info"),
	}
}

func (c Config) MySQLDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName, c.DBParams)
}

func get(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
