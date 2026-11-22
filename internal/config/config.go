package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config holds runtime configuration loaded from environment variables.
type Config struct {
	Host              string
	Port              int
	ReadHeaderTimeout time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
	MaxBodyBytes      int64
	ShutdownTimeout   time.Duration
	LogLevel          string

	// Storage configuration for upcoming persistence phase
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
}

// Load reads configuration from environment variables and applies validation.
func Load() (*Config, error) {
	cfg := &Config{
		Host:              getEnv("HTTP_HOST", "0.0.0.0"),
		Port:              getEnvInt("HTTP_PORT", 8080),
		ReadHeaderTimeout: getEnvDuration("READ_HEADER_TIMEOUT", 5*time.Second),
		ReadTimeout:       getEnvDuration("READ_TIMEOUT", 10*time.Second),
		WriteTimeout:      getEnvDuration("WRITE_TIMEOUT", 15*time.Second),
		IdleTimeout:       getEnvDuration("IDLE_TIMEOUT", 60*time.Second),
		MaxHeaderBytes:    getEnvInt("MAX_HEADER_BYTES", 1048576), // 1 MB
		MaxBodyBytes:      getEnvInt64("MAX_BODY_BYTES", 1048576), // 1 MB
		ShutdownTimeout:   getEnvDuration("SHUTDOWN_TIMEOUT", 10*time.Second),
		LogLevel:          strings.ToUpper(getEnv("LOG_LEVEL", "INFO")),

		DBHost:     getEnv("DB_HOST", "db"),
		DBPort:     getEnvInt("DB_PORT", 5432),
		DBName:     getEnv("DB_NAME", "filterkappi"),
		DBUser:     getEnv("DB_USER", "filterkappi"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return cfg, nil
}

// Validate checks that configuration values meet security and operational boundaries.
func (c *Config) Validate() error {
	if c.Port < 1 || c.Port > 65535 {
		return fmt.Errorf("HTTP_PORT must be between 1 and 65535, got %d", c.Port)
	}
	if c.ReadHeaderTimeout <= 0 {
		return errors.New("READ_HEADER_TIMEOUT must be positive")
	}
	if c.ReadTimeout <= 0 {
		return errors.New("READ_TIMEOUT must be positive")
	}
	if c.WriteTimeout <= 0 {
		return errors.New("WRITE_TIMEOUT must be positive")
	}
	if c.IdleTimeout <= 0 {
		return errors.New("IDLE_TIMEOUT must be positive")
	}
	if c.MaxHeaderBytes <= 0 {
		return errors.New("MAX_HEADER_BYTES must be positive")
	}
	if c.MaxBodyBytes <= 0 {
		return errors.New("MAX_BODY_BYTES must be positive")
	}
	if c.ShutdownTimeout <= 0 {
		return errors.New("SHUTDOWN_TIMEOUT must be positive")
	}

	switch c.LogLevel {
	case "DEBUG", "INFO", "WARN", "WARNING", "ERROR":
		// valid
	default:
		return fmt.Errorf("invalid LOG_LEVEL %q: must be DEBUG, INFO, WARN, or ERROR", c.LogLevel)
	}

	if c.DBPort < 1 || c.DBPort > 65535 {
		return fmt.Errorf("DB_PORT must be between 1 and 65535, got %d", c.DBPort)
	}

	return nil
}

func getEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	valStr, ok := os.LookupEnv(key)
	if !ok || valStr == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return -1 // Trigger validation failure
	}
	return val
}

func getEnvInt64(key string, defaultVal int64) int64 {
	valStr, ok := os.LookupEnv(key)
	if !ok || valStr == "" {
		return defaultVal
	}
	val, err := strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		return -1 // Trigger validation failure
	}
	return val
}

func getEnvDuration(key string, defaultVal time.Duration) time.Duration {
	valStr, ok := os.LookupEnv(key)
	if !ok || valStr == "" {
		return defaultVal
	}
	val, err := time.ParseDuration(valStr)
	if err != nil {
		return -1 // Trigger validation failure
	}
	return val
}
