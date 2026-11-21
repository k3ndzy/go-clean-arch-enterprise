package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Environment string

const (
	EnvDevelopment Environment = "development"
	EnvStaging     Environment = "staging"
	EnvProduction  Environment = "production"
)

type Config struct {
	App      AppConfig
	Server   ServerConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Name        string
	Environment Environment
	Version     string
}

type ServerConfig struct {
	Port            int
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

type DatabaseConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Name            string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

// Load tải cấu hình từ biến môi trường với giá trị mặc định an toàn
func Load() (*Config, error) {
	cfg := &Config{
		App: AppConfig{
			Name:        getEnv("APP_NAME", "go-clean-arch-enterprise"),
			Environment: Environment(getEnv("APP_ENV", "development")),
			Version:     getEnv("APP_VERSION", "v1.0.0"),
		},
		Server: ServerConfig{
			Port:            getEnvAsInt("SERVER_PORT", 8080),
			ReadTimeout:     getEnvAsDuration("SERVER_READ_TIMEOUT", 10*time.Second),
			WriteTimeout:    getEnvAsDuration("SERVER_WRITE_TIMEOUT", 15*time.Second),
			IdleTimeout:     getEnvAsDuration("SERVER_IDLE_TIMEOUT", 60*time.Second),
			ShutdownTimeout: getEnvAsDuration("SERVER_SHUTDOWN_TIMEOUT", 15*time.Second),
		},
		Database: DatabaseConfig{
			Host:            getEnv("DB_HOST", "localhost"),
			Port:            getEnvAsInt("DB_PORT", 5432),
			User:            getEnv("DB_USER", "postgres"),
			Password:        getEnv("DB_PASSWORD", "postgres"),
			Name:            getEnv("DB_NAME", "enterprise_db"),
			SSLMode:         getEnv("DB_SSL_MODE", "disable"),
			MaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 10),
			ConnMaxLifetime: getEnvAsDuration("DB_CONN_MAX_LIFETIME", 15*time.Minute),
		},
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("cấu hình không hợp lệ: %w", err)
	}

	return cfg, nil
}

func (c *Config) Validate() error {
	if c.Server.Port <= 0 || c.Server.Port > 65535 {
		return errors.New("cổng server phải nằm trong khoảng từ 1 đến 65535")
	}

	switch c.App.Environment {
	case EnvDevelopment, EnvStaging, EnvProduction:
	default:
		return fmt.Errorf("môi trường không hợp lệ: %s (chỉ chấp nhận development, staging, production)", c.App.Environment)
	}

	return nil
}

func getEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists && val != "" {
		return val
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	valStr := getEnv(key, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultVal
}

func getEnvAsDuration(key string, defaultVal time.Duration) time.Duration {
	valStr := getEnv(key, "")
	if val, err := time.ParseDuration(valStr); err == nil {
		return val
	}
	return defaultVal
}
