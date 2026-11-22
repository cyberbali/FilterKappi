package logging

import (
	"context"
	"log/slog"
	"os"
	"strings"
)

type contextKey string

const requestIDKey contextKey = "request_id"

// NewLogger initializes a structured JSON slog.Logger writing to stdout.
func NewLogger(level string) *slog.Logger {
	var slogLevel slog.Level
	switch strings.ToUpper(level) {
	case "DEBUG":
		slogLevel = slog.LevelDebug
	case "WARN", "WARNING":
		slogLevel = slog.LevelWarn
	case "ERROR":
		slogLevel = slog.LevelError
	default:
		slogLevel = slog.LevelInfo
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slogLevel,
	})

	return slog.New(handler)
}

// WithRequestID stores a request ID in the context.
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

// RequestIDFromContext extracts the request ID from context if present.
func RequestIDFromContext(ctx context.Context) string {
	if val, ok := ctx.Value(requestIDKey).(string); ok {
		return val
	}
	return ""
}
