// Package config loads process configuration from environment variables.
package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

// Config is process configuration loaded from environment variables.
type Config struct {
	HTTPAddr        string
	DatabaseURL     string
	Environment     string
	LogLevel        slog.Level
	ShutdownTimeout time.Duration
}

// Load reads configuration from the environment and applies safe defaults.
func Load() (Config, error) {
	cfg := Config{
		HTTPAddr:        env("HTTP_ADDR", ":8080"),
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		Environment:     env("APP_ENV", "development"),
		ShutdownTimeout: 10 * time.Second,
	}

	level, err := parseLogLevel(env("LOG_LEVEL", "info"))
	if err != nil {
		return Config{}, err
	}
	cfg.LogLevel = level

	return cfg, nil
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func parseLogLevel(raw string) (slog.Level, error) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(raw)); err != nil {
		return 0, fmt.Errorf("invalid LOG_LEVEL %q: %w", raw, err)
	}
	return level, nil
}
