package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestLoggingMiddleware tests that the logging middleware logs requests
func TestLoggingMiddleware(t *testing.T) {
	// Create a buffer to capture log output
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Create a simple handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Wrap the handler with the logging middleware
	wrapped := LoggingMiddleware(handler)

	// Create a test request
	req := httptest.NewRequest("GET", "/test", nil)
	req.RemoteAddr = "127.0.0.1:1234"

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Process the request
	wrapped.ServeHTTP(rr, req)

	// Check that the handler was called (status is OK)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	// Check that something was logged
	logOutput := buf.String()
	if !strings.Contains(logOutput, "GET /test 127.0.0.1:1234") {
		t.Errorf("Expected log to contain 'GET /test 127.0.0.1:1234', got: %s", logOutput)
	}
}

// TestAuthMiddleware_ValidKey tests that requests with valid API keys are processed
func TestAuthMiddleware_ValidKey(t *testing.T) {
	// Create a simple handler that returns a success message
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Authenticated"))
	})

	// Wrap the handler with the auth middleware
	wrapped := AuthMiddleware(handler)

	// Create a test request with a valid API key
	req := httptest.NewRequest("GET", "/secured", nil)
	req.Header.Set("X-API-Key", "valid-api-key")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Process the request
	wrapped.ServeHTTP(rr, req)

	// Check status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	// Check response body
	if rr.Body.String() != "Authenticated" {
		t.Errorf("Expected body 'Authenticated', got '%s'", rr.Body.String())
	}
}

// TestAuthMiddleware_InvalidKey tests that requests with invalid API keys are rejected
func TestAuthMiddleware_InvalidKey(t *testing.T) {
	// Create a simple handler (should not be called)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Handler was called with invalid API key")
	})

	// Wrap the handler with the auth middleware
	wrapped := AuthMiddleware(handler)

	// Create a test request with an invalid API key
	req := httptest.NewRequest("GET", "/secured", nil)
	req.Header.Set("X-API-Key", "invalid-key")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Process the request
	wrapped.ServeHTTP(rr, req)

	// Check status code (should be Unauthorized)
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, rr.Code)
	}

	// Check error message in response
	if !strings.Contains(rr.Body.String(), "Invalid API key") {
		t.Errorf("Expected body to contain 'Invalid API key', got '%s'", rr.Body.String())
	}
}

// TestRateLimitMiddleware tests that the rate limiting middleware restricts requests
func TestRateLimitMiddleware(t *testing.T) {
	// Create a simple handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Wrap the handler with rate limiting middleware (2 requests per minute)
	wrapped := RateLimitMiddleware(2)(handler)

	// Create a test request
	req := httptest.NewRequest("GET", "/rate-limited", nil)
	req.RemoteAddr = "127.0.0.1:1234" // Same IP for all requests

	// First request should succeed
	rr1 := httptest.NewRecorder()
	wrapped.ServeHTTP(rr1, req)
	if rr1.Code != http.StatusOK {
		t.Errorf("Expected first request to succeed with status %d, got %d", http.StatusOK, rr1.Code)
	}

	// Second request should succeed
	rr2 := httptest.NewRecorder()
	wrapped.ServeHTTP(rr2, req)
	if rr2.Code != http.StatusOK {
		t.Errorf("Expected second request to succeed with status %d, got %d", http.StatusOK, rr2.Code)
	}

	// Third request should be rate limited
	rr3 := httptest.NewRecorder()
	wrapped.ServeHTTP(rr3, req)
	if rr3.Code != http.StatusTooManyRequests {
		t.Errorf("Expected third request to be rate limited with status %d, got %d",
			http.StatusTooManyRequests, rr3.Code)
	}

	// Verify rate limit error message
	if !strings.Contains(rr3.Body.String(), "Rate limit exceeded") {
		t.Errorf("Expected body to contain 'Rate limit exceeded', got '%s'", rr3.Body.String())
	}
}

// TestRecoveryMiddleware tests that the recovery middleware catches panics
func TestRecoveryMiddleware(t *testing.T) {
	// Create a handler that panics
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("Test panic")
	})

	// Wrap the handler with recovery middleware
	wrapped := RecoveryMiddleware(handler)

	// Create a test request
	req := httptest.NewRequest("GET", "/panic", nil)

	// Create a response recorder
	rr := httptest.NewRecorder()

	// This should not panic due to the recovery middleware
	wrapped.ServeHTTP(rr, req)

	// Check status code (should be Internal Server Error)
	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rr.Code)
	}

	// Check error message
	if !strings.Contains(rr.Body.String(), "Internal Server Error") {
		t.Errorf("Expected body to contain 'Internal Server Error', got '%s'", rr.Body.String())
	}
}

// TestCORSMiddleware tests that CORS headers are added to responses
func TestCORSMiddleware(t *testing.T) {
	// Create a simple handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Wrap the handler with CORS middleware
	wrapped := CORSMiddleware(handler)

	// Create a test request
	req := httptest.NewRequest("GET", "/cors-test", nil)

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Process the request
	wrapped.ServeHTTP(rr, req)

	// Check CORS headers
	expectedHeaders := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
		"Access-Control-Allow-Headers": "Content-Type, Authorization",
	}

	for header, expectedValue := range expectedHeaders {
		if value := rr.Header().Get(header); value != expectedValue {
			t.Errorf("Expected header %s to be '%s', got '%s'", header, expectedValue, value)
		}
	}

	// Check that the handler was still called (body should be "OK")
	if rr.Body.String() != "OK" {
		t.Errorf("Expected body 'OK', got '%s'", rr.Body.String())
	}
}

// TestCORSMiddleware_Options tests that OPTIONS requests are handled correctly
func TestCORSMiddleware_Options(t *testing.T) {
	// Create a handler that should not be called for OPTIONS requests
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Handler was called for OPTIONS request")
	})

	// Wrap the handler with CORS middleware
	wrapped := CORSMiddleware(handler)

	// Create an OPTIONS request
	req := httptest.NewRequest("OPTIONS", "/cors-test", nil)

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Process the request
	wrapped.ServeHTTP(rr, req)

	// Check status code (should be OK)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	// Check that the response body is empty (handler not called)
	if rr.Body.String() != "" {
		t.Errorf("Expected empty body, got '%s'", rr.Body.String())
	}

	// Check CORS headers
	expectedHeaders := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
		"Access-Control-Allow-Headers": "Content-Type, Authorization",
	}

	for header, expectedValue := range expectedHeaders {
		if value := rr.Header().Get(header); value != expectedValue {
			t.Errorf("Expected header %s to be '%s', got '%s'", header, expectedValue, value)
		}
	}
}

// TestChain tests that middleware chaining works correctly
func TestChain(t *testing.T) {
	// Create test values to track middleware execution
	var (
		middleware1Called bool
		middleware2Called bool
		handlerCalled     bool
	)

	// Create two simple middleware functions
	middleware1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			middleware1Called = true
			next.ServeHTTP(w, r)
		})
	}

	middleware2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			middleware2Called = true
			next.ServeHTTP(w, r)
		})
	}

	// Create a simple handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerCalled = true
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Chain the middleware and handler
	chained := Chain(handler, middleware1, middleware2)

	// Create a test request
	req := httptest.NewRequest("GET", "/chain-test", nil)

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Process the request
	chained.ServeHTTP(rr, req)

	// Check that all middleware and the handler were called
	if !middleware1Called {
		t.Error("Middleware 1 was not called")
	}
	if !middleware2Called {
		t.Error("Middleware 2 was not called")
	}
	if !handlerCalled {
		t.Error("Handler was not called")
	}

	// Check response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
}

// TestFullServerSetup tests the entire server setup with middleware
func TestFullServerSetup(t *testing.T) {
	// Get the server with middleware
	mux := SetupMiddlewareServer()

	// Create a test server
	server := httptest.NewServer(mux)
	defer server.Close()

	// Test the hello endpoint with a valid API key
	req, err := http.NewRequest("GET", server.URL+"/hello", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("X-API-Key", "valid-api-key")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
