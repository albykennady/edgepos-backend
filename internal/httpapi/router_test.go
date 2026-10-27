package httpapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	router := NewRouter()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	// TODO 1: Assert that the status code is 200 OK.
	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)

	}

	// TODO 2: Assert that Content-Type is application/json.
	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Application expects application/json, got %q", got)

	}

	// TODO 3: Assert that the response body is {"status":"ok"}.
	if got := rec.Body.String(); got != `{"status":"ok"}` {
		t.Errorf("expected body %q, got %q", `{"status":"ok"}`, got)
	}
}
