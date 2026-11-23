package logging

import (
	"context"
	"testing"
)

func TestNewLogger(t *testing.T) {
	levels := []string{"DEBUG", "INFO", "WARN", "WARNING", "ERROR", "unknown"}
	for _, lvl := range levels {
		logger := NewLogger(lvl)
		if logger == nil {
			t.Errorf("expected non-nil logger for level %q", lvl)
		}
	}
}

func TestWithRequestID(t *testing.T) {
	ctx := context.Background()
	reqID := "req-123456"

	ctxWithID := WithRequestID(ctx, reqID)
	extracted := RequestIDFromContext(ctxWithID)

	if extracted != reqID {
		t.Errorf("expected request ID %q, got %q", reqID, extracted)
	}

	// Unset context should return empty string
	if empty := RequestIDFromContext(ctx); empty != "" {
		t.Errorf("expected empty string for context without request ID, got %q", empty)
	}
}
