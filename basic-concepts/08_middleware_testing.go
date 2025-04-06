package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Middleware is a function that wraps an http.Handler with additional functionality
type Middleware func(http.Handler) http.Handler

// LoggingMiddleware logs information about each request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log request information
		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}

// AuthMiddleware checks for a valid API key in the request header
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get API key from request header
		apiKey := r.Header.Get("X-API-Key")

		// Check if API key is valid (simplified example)
		if apiKey != "valid-api-key" {
			http.Error(w, "Unauthorized: Invalid API key", http.StatusUnauthorized)
			return
		}

		// API key is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

// RateLimitMiddleware limits the number of requests per client IP
func RateLimitMiddleware(requestsPerMinute int) Middleware {
	// In a real implementation, you'd use a more sophisticated tracking system
	// For this example, we'll use a simple map to track requests
	requestCounts := make(map[string]int)
	lastResetTime := time.Now()
	var mu = &sync.RWMutex{}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get client IP
			clientIP := r.RemoteAddr

			mu.RLock()
			// Check if we need to reset the counters
			if time.Since(lastResetTime) > time.Minute {
				mu.RUnlock()
				mu.Lock()
				// Reset counters if more than a minute has passed
				if time.Since(lastResetTime) > time.Minute {
					requestCounts = make(map[string]int)
					lastResetTime = time.Now()
				}
				mu.Unlock()
			} else {
				mu.RUnlock()
			}

			// Check if this client has exceeded the rate limit
			mu.Lock()
			requestCounts[clientIP]++
			count := requestCounts[clientIP]
			mu.Unlock()

			if count > requestsPerMinute {
				http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
				return
			}

			// Proceed to the next handler
			next.ServeHTTP(w, r)
		})
	}
}

// RecoveryMiddleware recovers from panics and responds with a 500 Internal Server Error
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Defer a function to recover from panics and return a 500 error
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// CORS middleware adds Cross-Origin Resource Sharing headers
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// Chain applies a series of middleware to a handler
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

// A simple handler to demonstrate middleware usage
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// PanicHandler intentionally panics to demonstrate the recovery middleware
func PanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("This is a deliberate panic!")
}

// Setup a server with middleware
func SetupMiddlewareServer() *http.ServeMux {
	mux := http.NewServeMux()

	// Setup routes with middleware
	mux.Handle("/hello", Chain(
		http.HandlerFunc(HelloHandler),
		LoggingMiddleware,
		AuthMiddleware,
		RateLimitMiddleware(10), // 10 requests per minute
		CORSMiddleware,
	))

	// Route with recovery middleware
	mux.Handle("/panic", Chain(
		http.HandlerFunc(PanicHandler),
		RecoveryMiddleware,
		LoggingMiddleware,
	))

	return mux
}

// StartMiddlewareServer starts the HTTP server with middleware
func StartMiddlewareServer() {
	mux := SetupMiddlewareServer()
	fmt.Println("Middleware server started on :8080")
	http.ListenAndServe(":8080", mux)
}

/*
Common Interview Questions about Middleware in Go:

1. What is middleware in the context of HTTP servers?
   - Middleware is a function that sits between the request and the handler
   - It intercepts HTTP requests/responses and can modify them or perform actions
   - It allows for separation of concerns and keeps handler code focused on business logic

2. How is middleware typically implemented in Go?
   - As a function that takes an http.Handler and returns a new http.Handler
   - It wraps the original handler and may execute code before and/or after calling it

3. What are common uses for middleware?
   - Logging, authentication, authorization, rate limiting, CORS, caching, etc.
   - Error handling and recovery from panics
   - Request/response modification (adding headers, compressing content)

4. How do you test middleware in Go?
   - Create simple handlers for testing
   - Use httptest package to create test requests and record responses
   - Verify that the middleware correctly modifies requests/responses

5. What is middleware chaining?
   - Applying multiple middleware layers in a specific order
   - The order matters: the first middleware in the chain is the outermost wrapper

6. What's the difference between middleware and handlers?
   - Handlers process requests and generate responses for specific routes
   - Middleware intercepts and processes all requests (or a subset) before/after handlers
   - Middleware is typically more generic and reusable
*/
