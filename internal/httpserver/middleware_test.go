package httpserver

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cyberbali/FilterKappi/internal/logging"
)

func TestRequestIDMiddleware_GeneratesNew(t *testing.T) {
	handler := RequestIDMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := logging.RequestIDFromContext(r.Context())
		if reqID == "" {
			t.Error("expected non-empty request ID in context")
		}
		w.WriteHeader(http.StatusOK)
	}))

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)

	handler.ServeHTTP(rec, req)

	headerReqID := rec.Header().Get("X-Request-ID")
	if headerReqID == "" {
		t.Fatal("expected X-Request-ID header in response")
	}
}

func TestRequestIDMiddleware_PreservesExisting(t *testing.T) {
	customID := "client-trace-12345"
	handler := RequestIDMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := logging.RequestIDFromContext(r.Context())
		if reqID != customID {
			t.Errorf("expected request ID %q, got %q", customID, reqID)
		}
		w.WriteHeader(http.StatusOK)
	}))

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("X-Request-ID", customID)

	handler.ServeHTTP(rec, req)

	if rec.Header().Get("X-Request-ID") != customID {
		t.Errorf("expected header %q, got %q", customID, rec.Header().Get("X-Request-ID"))
	}
}

func TestMaxBodyBytesMiddleware_LimitEnforced(t *testing.T) {
	const limit = 100 // 100 bytes limit

	handler := MaxBodyBytesMiddleware(limit)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := io.ReadAll(r.Body)
		if err != nil {
			if CheckMaxBytesError(err) {
				w.WriteHeader(http.StatusRequestEntityTooLarge)
				w.Write([]byte(`{"error":"request body too large"}`))
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))

	// Case 1: Body within limit
	t.Run("within limit", func(t *testing.T) {
		body := bytes.Repeat([]byte("a"), 50)
		req := httptest.NewRequest(http.MethodPost, "/upload", bytes.NewReader(body))
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("expected 200 OK, got %d", rec.Code)
		}
	})

	// Case 2: Body exceeds limit
	t.Run("exceeds limit", func(t *testing.T) {
		body := bytes.Repeat([]byte("a"), 200)
		req := httptest.NewRequest(http.MethodPost, "/upload", bytes.NewReader(body))
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusRequestEntityTooLarge {
			t.Errorf("expected 413 Request Entity Too Large, got %d", rec.Code)
		}
	})
}

func TestRecoveryMiddleware(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(io.Discard, nil))
	handler := RecoveryMiddleware(logger)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("unexpected runtime panic")
	}))

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/panic", nil)

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500 Internal Server Error, got %d", rec.Code)
	}

	expectedBody := `{"error":"internal server error"}`
	if rec.Body.String() != expectedBody {
		t.Errorf("expected generic error message %q, got %q", expectedBody, rec.Body.String())
	}
}
