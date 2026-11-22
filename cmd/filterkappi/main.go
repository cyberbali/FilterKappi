package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/cyberbali/FilterKappi/internal/config"
	"github.com/cyberbali/FilterKappi/internal/httpserver"
	"github.com/cyberbali/FilterKappi/internal/logging"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "startup failed: invalid configuration: %v\n", err)
		os.Exit(1)
	}

	logger := logging.NewLogger(cfg.LogLevel)

	logger.Info("initializing FilterKappi service",
		slog.String("host", cfg.Host),
		slog.Int("port", cfg.Port),
		slog.String("log_level", cfg.LogLevel),
	)

	server := httpserver.New(cfg, logger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	errChan := make(chan error, 1)
	go func() {
		errChan <- server.Start()
	}()

	select {
	case err := <-errChan:
		if err != nil {
			logger.Error("server encountered fatal error", slog.Any("error", err))
			os.Exit(1)
		}
	case <-ctx.Done():
		logger.Info("shutdown signal received, initiating graceful drain")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Error("graceful shutdown failed", slog.Any("error", err))
			os.Exit(1)
		}
		logger.Info("FilterKappi stopped gracefully")
	}
}
