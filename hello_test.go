package traefik_plugin_hello

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloMiddleware(t *testing.T) {
	cfg := &Config{HeaderKey: "X-Plugin", HeaderValue: "Hello"}
	h, err := New(context.Background(), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}), cfg, "test")
	if err != nil {
		t.Fatalf("New(): %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if got := rec.Header().Get("X-Plugin"); got != "Hello" {
		t.Fatalf("header = %q, want %q", got, "Hello")
	}
}
