package httpserver

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/cyberbali/FilterKappi/internal/config"
)

// Server wraps the standard library http.Server with lifecycle controls.
type Server struct {
	httpServer *http.Server
	logger     *slog.Logger
	cfg        *config.Config
}

// New constructs an HTTP server configured with explicit timeouts and security boundaries.
func New(cfg *config.Config, logger *slog.Logger) *Server {
	router := NewRouter(cfg, logger)

	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler:           router,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
		MaxHeaderBytes:    cfg.MaxHeaderBytes,
	}

	return &Server{
		httpServer: srv,
		logger:     logger,
		cfg:        cfg,
	}
}

// NewRouter sets up standard library ServeMux (Go 1.22+) and middleware chain.
func NewRouter(cfg *config.Config, logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()

	// Healthcheck endpoint (Go 1.22+ method pattern routing)
	mux.HandleFunc("GET /healthz", HealthzHandler)

	// Apply middleware stack
	var handler http.Handler = mux
	handler = MaxBodyBytesMiddleware(cfg.MaxBodyBytes)(handler)
	handler = LoggingMiddleware(logger)(handler)
	handler = RequestIDMiddleware(handler)
	handler = RecoveryMiddleware(logger)(handler)

	return handler
}

// HealthzHandler responds with a JSON status payload indicating service vitality.
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}

// Start begins listening for incoming requests.
func (s *Server) Start() error {
	s.logger.Info("starting HTTP server",
		slog.String("addr", s.httpServer.Addr),
		slog.Duration("read_header_timeout", s.httpServer.ReadHeaderTimeout),
		slog.Duration("read_timeout", s.httpServer.ReadTimeout),
		slog.Duration("write_timeout", s.httpServer.WriteTimeout),
		slog.Duration("idle_timeout", s.httpServer.IdleTimeout),
		slog.Int("max_header_bytes", s.httpServer.MaxHeaderBytes),
	)

	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("http server failure: %w", err)
	}
	return nil
}

// Shutdown gracefully stops the HTTP server.
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("shutting down HTTP server")
	return s.httpServer.Shutdown(ctx)
}

// UnderlyingServer exposes the underlying http.Server for inspection/testing.
func (s *Server) UnderlyingServer() *http.Server {
	return s.httpServer
}
