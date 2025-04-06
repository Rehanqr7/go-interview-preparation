package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Book represents book data
type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

// BookStore manages a collection of books with thread-safety
type BookStore struct {
	sync.RWMutex
	books     map[int]Book
	nextID    int
	idCounter int
}

// NewBookStore creates a new BookStore with some sample data
func NewBookStore() *BookStore {
	store := &BookStore{
		books:  make(map[int]Book),
		nextID: 1,
	}

	// Add some sample books
	store.AddBook(Book{
		Title:  "The Go Programming Language",
		Author: "Alan A. A. Donovan and Brian W. Kernighan",
		Price:  32.99,
	})

	store.AddBook(Book{
		Title:  "Concurrency in Go",
		Author: "Katherine Cox-Buday",
		Price:  34.99,
	})

	store.AddBook(Book{
		Title:  "Go in Action",
		Author: "William Kennedy",
		Price:  24.99,
	})

	return store
}

// GetBooks returns all books
func (bs *BookStore) GetBooks() []Book {
	bs.RLock()
	defer bs.RUnlock()

	books := make([]Book, 0, len(bs.books))
	for _, book := range bs.books {
		books = append(books, book)
	}
	return books
}

// GetBook retrieves a book by ID
func (bs *BookStore) GetBook(id int) (Book, bool) {
	bs.RLock()
	defer bs.RUnlock()

	book, exists := bs.books[id]
	return book, exists
}

// AddBook adds a new book and returns its ID
func (bs *BookStore) AddBook(book Book) int {
	bs.Lock()
	defer bs.Unlock()

	// Set ID and creation time
	book.ID = bs.nextID
	book.CreatedAt = time.Now()

	// Store book and increment ID counter
	bs.books[book.ID] = book
	bs.nextID++

	return book.ID
}

// UpdateBook updates an existing book
func (bs *BookStore) UpdateBook(id int, book Book) bool {
	bs.Lock()
	defer bs.Unlock()

	// Check if book exists
	_, exists := bs.books[id]
	if !exists {
		return false
	}

	// Preserve ID and creation time
	book.ID = id
	book.CreatedAt = bs.books[id].CreatedAt

	// Update book
	bs.books[id] = book
	return true
}

// DeleteBook removes a book by ID
func (bs *BookStore) DeleteBook(id int) bool {
	bs.Lock()
	defer bs.Unlock()

	_, exists := bs.books[id]
	if exists {
		delete(bs.books, id)
		return true
	}
	return false
}

// API handler functions

// handleGetBooks handles GET requests for all books
func handleGetBooks(w http.ResponseWriter, r *http.Request, store *BookStore) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	books := store.GetBooks()
	respondWithJSON(w, http.StatusOK, books)
}

// handleGetBook handles GET requests for a specific book
func handleGetBook(w http.ResponseWriter, r *http.Request, store *BookStore) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path
	// Expecting /books/{id}
	id, err := extractIDFromPath(r.URL.Path, "/books/")
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, exists := store.GetBook(id)
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, book)
}

// handleCreateBook handles POST requests to create a book
func handleCreateBook(w http.ResponseWriter, r *http.Request, store *BookStore) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate book data
	if book.Title == "" || book.Author == "" || book.Price <= 0 {
		http.Error(w, "Invalid book data: title, author and price are required", http.StatusBadRequest)
		return
	}

	// Add book to store
	id := store.AddBook(book)

	// Return the created book with its ID
	createdBook, _ := store.GetBook(id)
	respondWithJSON(w, http.StatusCreated, createdBook)
}

// handleUpdateBook handles PUT requests to update a book
func handleUpdateBook(w http.ResponseWriter, r *http.Request, store *BookStore) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path
	id, err := extractIDFromPath(r.URL.Path, "/books/")
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Parse request body
	var book Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate book data
	if book.Title == "" || book.Author == "" || book.Price <= 0 {
		http.Error(w, "Invalid book data: title, author and price are required", http.StatusBadRequest)
		return
	}

	// Update book
	success := store.UpdateBook(id, book)
	if !success {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Return the updated book
	updatedBook, _ := store.GetBook(id)
	respondWithJSON(w, http.StatusOK, updatedBook)
}

// handleDeleteBook handles DELETE requests to delete a book
func handleDeleteBook(w http.ResponseWriter, r *http.Request, store *BookStore) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path
	id, err := extractIDFromPath(r.URL.Path, "/books/")
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Delete book
	success := store.DeleteBook(id)
	if !success {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusNoContent)
}

// Utility functions

// respondWithJSON writes a JSON response
func respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// extractIDFromPath extracts and validates ID from URL path
func extractIDFromPath(path, prefix string) (int, error) {
	// Remove prefix from path
	idStr := path[len(prefix):]

	// Convert to integer
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("invalid ID: %s", idStr)
	}

	return id, nil
}

// Define a middleware type
type Middleware func(http.HandlerFunc) http.HandlerFunc

// loggingMiddleware logs request information
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(startTime))
	}
}

// applyMiddleware applies middlewares to a handler function
func applyMiddleware(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func main() {
	// Create book store
	store := NewBookStore()

	// Create router
	mux := http.NewServeMux()

	// Register routes with middleware
	mux.HandleFunc("/books", applyMiddleware(
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				handleGetBooks(w, r, store)
			case http.MethodPost:
				handleCreateBook(w, r, store)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		},
		loggingMiddleware,
	))

	mux.HandleFunc("/books/", applyMiddleware(
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				handleGetBook(w, r, store)
			case http.MethodPut:
				handleUpdateBook(w, r, store)
			case http.MethodDelete:
				handleDeleteBook(w, r, store)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		},
		loggingMiddleware,
	))

	// Start server
	port := ":8080"
	fmt.Printf("Starting RESTful API server on http://localhost%s\n", port)
	fmt.Println("API Endpoints:")
	fmt.Println("  GET    /books      - List all books")
	fmt.Println("  GET    /books/{id} - Get a specific book")
	fmt.Println("  POST   /books      - Create a new book")
	fmt.Println("  PUT    /books/{id} - Update a book")
	fmt.Println("  DELETE /books/{id} - Delete a book")

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

/*
This project demonstrates:

1. RESTful API design in Go
   - Resource-based URL structure
   - Appropriate HTTP methods (GET, POST, PUT, DELETE)
   - HTTP status codes
   - JSON responses

2. Concurrency-safe data access
   - Using RWMutex to protect a shared data store
   - Read locks for GET operations
   - Write locks for POST, PUT, DELETE operations

3. HTTP server implementation
   - Request routing
   - Request/response handling
   - Middleware pattern

4. Common Go patterns
   - Middleware chaining
   - Handler functions
   - Error handling

5. JSON serialization/deserialization
   - Using struct tags to control JSON field names
   - Request body parsing
   - Response generation

To test, run this server and use curl or a tool like Postman to make API requests:

# List all books
curl -X GET http://localhost:8080/books

# Get a specific book
curl -X GET http://localhost:8080/books/1

# Create a new book
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title":"Learning Go","author":"Jon Bodner","price":29.99}'

# Update a book
curl -X PUT http://localhost:8080/books/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"The Go Programming Language","author":"Donovan & Kernighan","price":39.99}'

# Delete a book
curl -X DELETE http://localhost:8080/books/1

*/
