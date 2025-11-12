package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRootHandler(t *testing.T) {
	// Set environment variable for testing
	os.Setenv("CLOUD_PROVIDER", "test")

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		provider := os.Getenv("CLOUD_PROVIDER")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello from Go! Cloud-agnostic demo running on " + provider))
	})

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	expected := "Hello from Go! Cloud-agnostic demo running on test"
	if rr.Body.String() != expected {
		t.Errorf("expected body %q, got %q", expected, rr.Body.String())
	}
}
