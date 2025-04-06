package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Response represents a simple API response
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// UserHandler handles user operations
type UserHandler struct {
	users map[string]User
}

// NewUserHandler creates a new user handler
func NewUserHandler() *UserHandler {
	return &UserHandler{
		users: map[string]User{
			"1": {
				ID:        1,
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john@example.com",
				Age:       30,
			},
		},
	}
}

// GetUser handles GET requests for users
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get user ID from URL query parameter
	userID := r.URL.Query().Get("id")
	if userID == "" {
		respondWithJSON(w, http.StatusBadRequest, Response{
			Status: "error",
			Error:  "user ID is required",
		})
		return
	}

	// Lookup user
	user, exists := h.users[userID]
	if !exists {
		respondWithJSON(w, http.StatusNotFound, Response{
			Status: "error",
			Error:  "user not found",
		})
		return
	}

	// Return user data
	respondWithJSON(w, http.StatusOK, Response{
		Status: "success",
		Data:   user,
	})
}

// CreateUser handles POST requests to create users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var user User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithJSON(w, http.StatusBadRequest, Response{
			Status: "error",
			Error:  "invalid request payload",
		})
		return
	}
	defer r.Body.Close()

	// Validate user
	if err := ValidateUser(user); err != nil {
		respondWithJSON(w, http.StatusBadRequest, Response{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	// Store user (in a real app, we'd generate a unique ID)
	userID := fmt.Sprintf("%d", user.ID)
	h.users[userID] = user

	// Return success
	respondWithJSON(w, http.StatusCreated, Response{
		Status:  "success",
		Message: "User created successfully",
		Data:    user,
	})
}

// Helper function to respond with JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// SetupRoutes configures the HTTP routes
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	userHandler := NewUserHandler()

	mux.HandleFunc("/user", userHandler.GetUser)
	mux.HandleFunc("/user/create", userHandler.CreateUser)

	return mux
}

// StartServer starts the HTTP server
func StartServer() {
	mux := SetupRoutes()
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", mux)
}

// Demo function to show how to use the server
func demonstrateHTTPServer() {
	fmt.Println("To start the HTTP server, call StartServer()")
	fmt.Println("Example API endpoints:")
	fmt.Println("  GET /user?id=1 - Get user with ID 1")
	fmt.Println("  POST /user/create - Create a new user with JSON payload")
}

/*
Common Interview Questions about HTTP Testing in Go:

1. How do you test HTTP handlers in Go?
   - Use the net/http/httptest package to create a test server and send requests to it
   - Use the httptest.NewRecorder() to capture the response
   - Create a request with httptest.NewRequest() and pass it to the handler

2. What is httptest.ResponseRecorder?
   - It's a mock http.ResponseWriter used to record the response status, headers, and body
   - It implements the http.ResponseWriter interface but records the data instead of writing to a connection

3. How do you test middleware in Go?
   - Create a simple handler that the middleware wraps
   - Test that the middleware correctly modifies requests/responses

4. What are good practices for HTTP testing?
   - Use table-driven tests for different scenarios
   - Test edge cases and error handling
   - Isolate dependencies using interfaces and mocks
   - Use subtests for different test cases

5. How do you mock external API calls in tests?
   - Create interfaces for external services
   - Implement mock versions of these interfaces for testing
   - Use dependency injection to provide mock implementations

6. How do you test JSON serialization/deserialization?
   - Create test cases with known inputs and expected outputs
   - Deserialize JSON responses and compare with expected structures
   - Test error cases with malformed JSON
*/
