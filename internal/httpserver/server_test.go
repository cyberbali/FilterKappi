package httpserver

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cyberbali/FilterKappi/internal/config"
)

func newTestConfig() *config.Config {
	return &config.Config{
		Host:              "127.0.0.1",
		Port:              8080,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1048576,
		MaxBodyBytes:      1048576,
		ShutdownTimeout:   5 * time.Second,
		LogLevel:          "INFO",
	}
}

func TestHealthzEndpoint(t *testing.T) {
	cfg := newTestConfig()
	logger := slog.New(slog.NewJSONHandler(io.Discard, nil))
	router := NewRouter(cfg, logger)

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	contentType := rec.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %q", contentType)
	}

	reqID := rec.Header().Get("X-Request-ID")
	if reqID == "" {
		t.Error("expected X-Request-ID header to be present")
	}

	expectedBody := `{"status":"ok"}`
	if rec.Body.String() != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, rec.Body.String())
	}
}

func TestHealthz_MethodNotAllowed(t *testing.T) {
	cfg := newTestConfig()
	logger := slog.New(slog.NewJSONHandler(io.Discard, nil))
	router := NewRouter(cfg, logger)

	// POST to GET-only route in Go 1.22+ ServeMux
	req := httptest.NewRequest(http.MethodPost, "/healthz", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405 Method Not Allowed, got %d", rec.Code)
	}
}

func TestServerTimeoutsConfigured(t *testing.T) {
	cfg := newTestConfig()
	logger := slog.New(slog.NewJSONHandler(io.Discard, nil))
	srv := New(cfg, logger)

	raw := srv.UnderlyingServer()
	if raw.ReadHeaderTimeout != cfg.ReadHeaderTimeout {
		t.Errorf("expected ReadHeaderTimeout %v, got %v", cfg.ReadHeaderTimeout, raw.ReadHeaderTimeout)
	}
	if raw.ReadTimeout != cfg.ReadTimeout {
		t.Errorf("expected ReadTimeout %v, got %v", cfg.ReadTimeout, raw.ReadTimeout)
	}
	if raw.WriteTimeout != cfg.WriteTimeout {
		t.Errorf("expected WriteTimeout %v, got %v", cfg.WriteTimeout, raw.WriteTimeout)
	}
	if raw.IdleTimeout != cfg.IdleTimeout {
		t.Errorf("expected IdleTimeout %v, got %v", cfg.IdleTimeout, raw.IdleTimeout)
	}
	if raw.MaxHeaderBytes != cfg.MaxHeaderBytes {
		t.Errorf("expected MaxHeaderBytes %d, got %d", cfg.MaxHeaderBytes, raw.MaxHeaderBytes)
	}
}

func TestServerShutdown(t *testing.T) {
	cfg := newTestConfig()
	cfg.Port = 0 // Ephemeral port
	logger := slog.New(slog.NewJSONHandler(io.Discard, nil))
	srv := New(cfg, logger)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		t.Fatalf("expected clean shutdown, got: %v", err)
	}
}
