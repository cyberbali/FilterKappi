package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cyberbali/FilterKappi/internal/config"
	"github.com/cyberbali/FilterKappi/internal/httpserver"
)

func TestIntegration_Healthz(t *testing.T) {
	cfg := &config.Config{
		Host:              "127.0.0.1",
		Port:              0,
		ReadHeaderTimeout: 2 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		MaxHeaderBytes:    1048576,
		MaxBodyBytes:      1048576,
		ShutdownTimeout:   2 * time.Second,
		LogLevel:          "INFO",
	}

	logger := slog.New(slog.NewJSONHandler(io.Discard, nil))
	router := httpserver.NewRouter(cfg, logger)

	ts := httptest.NewServer(router)
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/healthz", ts.URL))
	if err != nil {
		t.Fatalf("failed to query /healthz: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got: %d", resp.StatusCode)
	}

	if ct := resp.Header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected Content-Type application/json, got: %q", ct)
	}

	if reqID := resp.Header.Get("X-Request-ID"); reqID == "" {
		t.Error("expected non-empty X-Request-ID response header")
	}

	var body map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode json body: %v", err)
	}

	if body["status"] != "ok" {
		t.Errorf("expected status 'ok', got: %q", body["status"])
	}
}

func TestIntegration_MaxBodyLimit(t *testing.T) {
	const smallLimit = 50

	cfg := &config.Config{
		Host:              "127.0.0.1",
		Port:              0,
		ReadHeaderTimeout: 2 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		MaxHeaderBytes:    1048576,
		MaxBodyBytes:      smallLimit,
		ShutdownTimeout:   2 * time.Second,
		LogLevel:          "INFO",
	}

	// Handler that attempts to read the entire body
	mux := http.NewServeMux()
	mux.HandleFunc("POST /echo", func(w http.ResponseWriter, r *http.Request) {
		_, err := io.ReadAll(r.Body)
		if err != nil {
			if httpserver.CheckMaxBytesError(err) {
				w.WriteHeader(http.StatusRequestEntityTooLarge)
				w.Write([]byte(`{"error":"payload too large"}`))
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	var handler http.Handler = mux
	handler = httpserver.MaxBodyBytesMiddleware(cfg.MaxBodyBytes)(handler)
	handler = httpserver.RequestIDMiddleware(handler)

	ts := httptest.NewServer(handler)
	defer ts.Close()

	// Exceeds 50 bytes limit
	largePayload := bytes.Repeat([]byte("X"), 100)
	resp, err := http.Post(fmt.Sprintf("%s/echo", ts.URL), "text/plain", bytes.NewReader(largePayload))
	if err != nil {
		t.Fatalf("failed to send post request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusRequestEntityTooLarge {
		t.Fatalf("expected status 413, got: %d", resp.StatusCode)
	}
}
