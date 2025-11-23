package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config содержит конфигурацию приложения
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Log      LogConfig
}

// ServerConfig конфигурация сервера
type ServerConfig struct {
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// DatabaseConfig конфигурация БД
type DatabaseConfig struct {
	URL            string
	MaxOpenConns   int
	MaxIdleConns   int
	ConnMaxLifeMS  int
	MigrationsPath string
}

// LogConfig конфигурация логирования
type LogConfig struct {
	Level string // debug, info, warn, error
}

// Load загружает конфигурацию из переменных окружения
func Load() Config {
	return Config{
		Server: ServerConfig{
			Host:         getEnv("SERVER_HOST", "0.0.0.0"),
			Port:         getEnvInt("SERVER_PORT", 8080),
			ReadTimeout:  getEnvDuration("SERVER_READ_TIMEOUT", 15*time.Second),
			WriteTimeout: getEnvDuration("SERVER_WRITE_TIMEOUT", 15*time.Second),
			IdleTimeout:  getEnvDuration("SERVER_IDLE_TIMEOUT", 60*time.Second),
		},
		Database: DatabaseConfig{
			URL:            getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/go_cmms"),
			MaxOpenConns:   getEnvInt("DATABASE_MAX_OPEN_CONNS", 25),
			MaxIdleConns:   getEnvInt("DATABASE_MAX_IDLE_CONNS", 5),
			ConnMaxLifeMS:  getEnvInt("DATABASE_CONN_MAX_LIFE_MS", 5*60*1000),
			MigrationsPath: getEnv("DATABASE_MIGRATIONS_PATH", "file://internal/infrastructure/persistence/sqlc/migrations"),
		},
		Log: LogConfig{
			Level: getEnv("LOG_LEVEL", "info"),
		},
	}
}

// Address возвращает полный адрес сервера
func (sc ServerConfig) Address() string {
	return fmt.Sprintf("%s:%d", sc.Host, sc.Port)
}

// getEnv читает переменную окружения с значением по умолчанию
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvInt читает переменную окружения как целое число
func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

// getEnvDuration читает переменную окружения как длительность
func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}
