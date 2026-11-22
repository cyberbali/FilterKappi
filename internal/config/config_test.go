package config

import (
	"os"
	"testing"
	"time"
)

func clearEnv() {
	envKeys := []string{
		"HTTP_HOST", "HTTP_PORT", "READ_HEADER_TIMEOUT", "READ_TIMEOUT",
		"WRITE_TIMEOUT", "IDLE_TIMEOUT", "MAX_HEADER_BYTES", "MAX_BODY_BYTES",
		"SHUTDOWN_TIMEOUT", "LOG_LEVEL", "DB_HOST", "DB_PORT", "DB_NAME",
		"DB_USER", "DB_PASSWORD",
	}
	for _, key := range envKeys {
		os.Unsetenv(key)
	}
}

func TestLoad_Defaults(t *testing.T) {
	clearEnv()

	cfg, err := Load()
	if err != nil {
		t.Fatalf("expected Load to succeed with defaults, got: %v", err)
	}

	if cfg.Host != "0.0.0.0" {
		t.Errorf("expected default host 0.0.0.0, got %q", cfg.Host)
	}
	if cfg.Port != 8080 {
		t.Errorf("expected default port 8080, got %d", cfg.Port)
	}
	if cfg.ReadHeaderTimeout != 5*time.Second {
		t.Errorf("expected ReadHeaderTimeout 5s, got %v", cfg.ReadHeaderTimeout)
	}
	if cfg.ReadTimeout != 10*time.Second {
		t.Errorf("expected ReadTimeout 10s, got %v", cfg.ReadTimeout)
	}
	if cfg.WriteTimeout != 15*time.Second {
		t.Errorf("expected WriteTimeout 15s, got %v", cfg.WriteTimeout)
	}
	if cfg.IdleTimeout != 60*time.Second {
		t.Errorf("expected IdleTimeout 60s, got %v", cfg.IdleTimeout)
	}
	if cfg.MaxHeaderBytes != 1048576 {
		t.Errorf("expected MaxHeaderBytes 1048576, got %d", cfg.MaxHeaderBytes)
	}
	if cfg.MaxBodyBytes != 1048576 {
		t.Errorf("expected MaxBodyBytes 1048576, got %d", cfg.MaxBodyBytes)
	}
	if cfg.ShutdownTimeout != 10*time.Second {
		t.Errorf("expected ShutdownTimeout 10s, got %v", cfg.ShutdownTimeout)
	}
	if cfg.LogLevel != "INFO" {
		t.Errorf("expected LogLevel INFO, got %q", cfg.LogLevel)
	}
}

func TestLoad_CustomEnv(t *testing.T) {
	clearEnv()
	t.Setenv("HTTP_HOST", "127.0.0.1")
	t.Setenv("HTTP_PORT", "9090")
	t.Setenv("READ_HEADER_TIMEOUT", "3s")
	t.Setenv("READ_TIMEOUT", "7s")
	t.Setenv("WRITE_TIMEOUT", "12s")
	t.Setenv("IDLE_TIMEOUT", "45s")
	t.Setenv("MAX_HEADER_BYTES", "2048")
	t.Setenv("MAX_BODY_BYTES", "4096")
	t.Setenv("SHUTDOWN_TIMEOUT", "15s")
	t.Setenv("LOG_LEVEL", "debug")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("expected Load to succeed with custom env, got: %v", err)
	}

	if cfg.Host != "127.0.0.1" || cfg.Port != 9090 {
		t.Errorf("unexpected host/port: %s:%d", cfg.Host, cfg.Port)
	}
	if cfg.ReadHeaderTimeout != 3*time.Second {
		t.Errorf("unexpected ReadHeaderTimeout: %v", cfg.ReadHeaderTimeout)
	}
	if cfg.MaxBodyBytes != 4096 {
		t.Errorf("unexpected MaxBodyBytes: %d", cfg.MaxBodyBytes)
	}
	if cfg.LogLevel != "DEBUG" {
		t.Errorf("unexpected LogLevel: %q", cfg.LogLevel)
	}
}

func TestValidate_InvalidPort(t *testing.T) {
	testCases := []struct {
		name string
		port string
	}{
		{"non-numeric", "invalid"},
		{"zero", "0"},
		{"negative", "-8080"},
		{"too high", "70000"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			clearEnv()
			t.Setenv("HTTP_PORT", tc.port)
			_, err := Load()
			if err == nil {
				t.Fatalf("expected error for HTTP_PORT=%q, got nil", tc.port)
			}
		})
	}
}

func TestValidate_NegativeTimeouts(t *testing.T) {
	testCases := []string{
		"READ_HEADER_TIMEOUT",
		"READ_TIMEOUT",
		"WRITE_TIMEOUT",
		"IDLE_TIMEOUT",
		"SHUTDOWN_TIMEOUT",
	}

	for _, envKey := range testCases {
		t.Run(envKey, func(t *testing.T) {
			clearEnv()
			t.Setenv(envKey, "-5s")
			_, err := Load()
			if err == nil {
				t.Fatalf("expected error for negative %s, got nil", envKey)
			}
		})
	}
}

func TestValidate_InvalidMaxBodyBytes(t *testing.T) {
	clearEnv()
	t.Setenv("MAX_BODY_BYTES", "0")
	if _, err := Load(); err == nil {
		t.Fatal("expected error for MAX_BODY_BYTES=0, got nil")
	}

	clearEnv()
	t.Setenv("MAX_BODY_BYTES", "-100")
	if _, err := Load(); err == nil {
		t.Fatal("expected error for negative MAX_BODY_BYTES, got nil")
	}
}

func TestValidate_InvalidLogLevel(t *testing.T) {
	clearEnv()
	t.Setenv("LOG_LEVEL", "TRACE_VERBOSE")
	if _, err := Load(); err == nil {
		t.Fatal("expected error for unsupported LOG_LEVEL, got nil")
	}
}
