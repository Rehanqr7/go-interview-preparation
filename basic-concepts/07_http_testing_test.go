package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test GetUser handler with a valid user ID
func TestGetUser_ValidID(t *testing.T) {
	// Create a new UserHandler with the test data
	handler := NewUserHandler()

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(handler.GetUser))
	defer server.Close()

	// Make a GET request to the server with a valid user ID
	resp, err := http.Get(server.URL + "?id=1")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Decode the response
	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Check response status
	if response.Status != "success" {
		t.Errorf("Expected status 'success', got '%s'", response.Status)
	}

	// Check response data
	userData, ok := response.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("Data is not a map[string]interface{}")
	}

	expectedFirstName := "John"
	if userData["FirstName"] != expectedFirstName {
		t.Errorf("Expected FirstName '%s', got '%v'", expectedFirstName, userData["FirstName"])
	}
}

// Test GetUser handler with an invalid user ID
func TestGetUser_InvalidID(t *testing.T) {
	// Create a new UserHandler with the test data
	handler := NewUserHandler()

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(handler.GetUser))
	defer server.Close()

	// Make a GET request to the server with an invalid user ID
	resp, err := http.Get(server.URL + "?id=999")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, resp.StatusCode)
	}

	// Decode the response
	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Check response status
	if response.Status != "error" {
		t.Errorf("Expected status 'error', got '%s'", response.Status)
	}

	// Check error message
	expectedError := "user not found"
	if response.Error != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, response.Error)
	}
}

// Test GetUser handler with no user ID
func TestGetUser_MissingID(t *testing.T) {
	// Create a new UserHandler with the test data
	handler := NewUserHandler()

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(handler.GetUser))
	defer server.Close()

	// Make a GET request to the server with no user ID
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	// Decode the response
	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Check response status
	if response.Status != "error" {
		t.Errorf("Expected status 'error', got '%s'", response.Status)
	}

	// Check error message
	expectedError := "user ID is required"
	if response.Error != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, response.Error)
	}
}

// Test CreateUser handler with valid user data
func TestCreateUser_ValidData(t *testing.T) {
	// Create a new UserHandler with the test data
	handler := NewUserHandler()

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(handler.CreateUser))
	defer server.Close()

	// Create a user to send to the server
	newUser := User{
		ID:        2,
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@example.com",
		Age:       25,
	}

	// Convert user to JSON
	userData, err := json.Marshal(newUser)
	if err != nil {
		t.Fatalf("Failed to marshal user: %v", err)
	}

	// Create a POST request with the user data
	resp, err := http.Post(server.URL, "application/json", bytes.NewBuffer(userData))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	// Decode the response
	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Check response status
	if response.Status != "success" {
		t.Errorf("Expected status 'success', got '%s'", response.Status)
	}

	// Check response message
	expectedMessage := "User created successfully"
	if response.Message != expectedMessage {
		t.Errorf("Expected message '%s', got '%s'", expectedMessage, response.Message)
	}

	// Verify the user was added to the handler's map
	if user, exists := handler.users["2"]; !exists {
		t.Errorf("User was not added to the users map")
	} else if user.FirstName != "Jane" {
		t.Errorf("Expected FirstName 'Jane', got '%s'", user.FirstName)
	}
}

// Test CreateUser handler with invalid user data
func TestCreateUser_InvalidData(t *testing.T) {
	// Create a new UserHandler with the test data
	handler := NewUserHandler()

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(handler.CreateUser))
	defer server.Close()

	// Create an invalid user (missing FirstName)
	invalidUser := User{
		ID:        3,
		FirstName: "", // Invalid: empty first name
		LastName:  "Doe",
		Email:     "invalid@example.com",
		Age:       30,
	}

	// Convert user to JSON
	userData, err := json.Marshal(invalidUser)
	if err != nil {
		t.Fatalf("Failed to marshal user: %v", err)
	}

	// Create a POST request with the invalid user data
	resp, err := http.Post(server.URL, "application/json", bytes.NewBuffer(userData))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	// Decode the response
	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Check response status
	if response.Status != "error" {
		t.Errorf("Expected status 'error', got '%s'", response.Status)
	}

	// Check error message contains "first name"
	if response.Error == "" || response.Error != "first name cannot be empty" {
		t.Errorf("Expected error about empty first name, got '%s'", response.Error)
	}
}

// Test using httptest.ResponseRecorder directly
func TestGetUser_WithResponseRecorder(t *testing.T) {
	// Create a new UserHandler with the test data
	handler := NewUserHandler()

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/user?id=1", nil)

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler.GetUser(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	// Check the Content-Type header
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Expected Content-Type %s, got %s", expectedContentType, contentType)
	}

	// Decode the response
	var response Response
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Check response status
	if response.Status != "success" {
		t.Errorf("Expected status 'success', got '%s'", response.Status)
	}
}

// Test the entire HTTP router/server
func TestRouter(t *testing.T) {
	// Create the router from the SetupRoutes function
	router := SetupRoutes()

	// Create a test server with the router
	server := httptest.NewServer(router)
	defer server.Close()

	// Make a GET request to the /user endpoint
	resp, err := http.Get(server.URL + "/user?id=1")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Try another endpoint with a POST request to make sure the router works
	newUser := User{
		ID:        4,
		FirstName: "Bob",
		LastName:  "Smith",
		Email:     "bob@example.com",
		Age:       40,
	}

	// Convert user to JSON
	userData, err := json.Marshal(newUser)
	if err != nil {
		t.Fatalf("Failed to marshal user: %v", err)
	}

	// Create a POST request to the /user/create endpoint
	resp2, err := http.Post(server.URL+"/user/create", "application/json", bytes.NewBuffer(userData))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp2.Body.Close()

	// Check the response status code
	if resp2.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp2.StatusCode)
	}
}
