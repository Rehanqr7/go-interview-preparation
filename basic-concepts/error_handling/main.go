package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// BASIC ERROR HANDLING

// Simple function that returns an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Function that uses fmt.Errorf to create formatted error
func validateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("age %d is negative", age)
	}
	if age > 150 {
		return fmt.Errorf("age %d is too high", age)
	}
	return nil
}

// CUSTOM ERROR TYPES

// Define a custom error type
type InputValidationError struct {
	Field string
	Msg   string
}

// Implement the error interface
func (e InputValidationError) Error() string {
	return fmt.Sprintf("validation error: %s %s", e.Field, e.Msg)
}

// Function that returns a custom error
func validateNameInput(name string) error {
	if name == "" {
		return InputValidationError{Field: "name", Msg: "cannot be empty"}
	}
	if len(name) < 2 {
		return InputValidationError{Field: "name", Msg: "too short"}
	}
	if len(name) > 50 {
		return InputValidationError{Field: "name", Msg: "too long"}
	}
	return nil
}

// MULTIPLE ERROR TYPES

// Different error types for different error conditions
type SyntaxError struct {
	Line int
	Msg  string
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf("syntax error at line %d: %s", e.Line, e.Msg)
}

type RuntimeError struct {
	Time  string
	Msg   string
	Fatal bool
}

func (e RuntimeError) Error() string {
	if e.Fatal {
		return fmt.Sprintf("fatal runtime error at %s: %s", e.Time, e.Msg)
	}
	return fmt.Sprintf("runtime error at %s: %s", e.Time, e.Msg)
}

// ERROR WRAPPING (Go 1.13+)

// Function that wraps errors
func getFileContents(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		// Wrap the error with additional context
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(data), nil
}

// SENTINEL ERRORS

// Predefined errors for specific error conditions
var (
	ErrNotFound      = errors.New("item not found")
	ErrInvalidInput  = errors.New("invalid input")
	ErrUnauthorized  = errors.New("unauthorized access")
	ErrInternalError = errors.New("internal server error")
)

// Function that returns sentinel errors
func findItem(id string) (string, error) {
	if id == "" {
		return "", ErrInvalidInput
	}
	// Simulate no item found
	return "", ErrNotFound
}

// ERROR CHECKING PATTERNS

// Parse a positive integer with error handling
func parsePositiveInt(s string) (int, error) {
	// Convert string to int
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse integer: %w", err)
	}

	// Validate if it's positive
	if num <= 0 {
		return 0, errors.New("integer must be positive")
	}

	return num, nil
}

// PANIC AND RECOVER

// Function that panics
func mustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("failed to parse '%s' as integer: %v", s, err))
	}
	return n
}

// Function that uses recover to handle panics
func safeParse(s string) (n int, err error) {
	// Set up a deferred function to recover from panic
	defer func() {
		// Recover from panic if one occurs
		if r := recover(); r != nil {
			// Convert panic to error
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()

	// This might panic
	n = mustParseInt(s)
	return n, nil
}

func main() {
	fmt.Println("=== BASIC ERROR HANDLING ===")

	// Basic error handling
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Handling formatted errors
	err = validateAge(25)
	if err != nil {
		fmt.Println("Age validation error:", err)
	} else {
		fmt.Println("Age is valid")
	}

	err = validateAge(200)
	if err != nil {
		fmt.Println("Age validation error:", err)
	}

	fmt.Println("\n=== CUSTOM ERROR TYPES ===")

	// Custom error types
	err = validateNameInput("John")
	if err != nil {
		fmt.Println("Name validation error:", err)
	} else {
		fmt.Println("Name is valid")
	}

	err = validateNameInput("")
	if err != nil {
		fmt.Println("Name validation error:", err)

		// Type assertion
		if valErr, ok := err.(InputValidationError); ok {
			fmt.Printf("Field '%s' has error: %s\n", valErr.Field, valErr.Msg)
		}
	}

	fmt.Println("\n=== TYPE ASSERTION AND TYPE SWITCH ===")

	// Create different error types
	var err1 error = SyntaxError{Line: 42, Msg: "unexpected semicolon"}
	var err2 error = RuntimeError{Time: "2023-04-01 15:30:00", Msg: "out of memory", Fatal: true}

	// Handling different error types with type assertion
	if syntaxErr, ok := err1.(SyntaxError); ok {
		fmt.Printf("Syntax error on line %d: %s\n", syntaxErr.Line, syntaxErr.Msg)
	}

	// Handling different error types with type switch
	switch e := err2.(type) {
	case SyntaxError:
		fmt.Printf("Syntax error on line %d: %s\n", e.Line, e.Msg)
	case RuntimeError:
		fmt.Printf("Runtime error (%v) at %s: %s\n", e.Fatal, e.Time, e.Msg)
	default:
		fmt.Printf("Unknown error: %v\n", e)
	}

	fmt.Println("\n=== ERROR WRAPPING ===")

	// Error wrapping
	_, err = getFileContents("nonexistent-file.txt")
	if err != nil {
		fmt.Println("Error:", err)

		// Unwrap the error (Go 1.13+)
		fmt.Println("Unwrapped error:", errors.Unwrap(err))

		// Check if an error is wrapped inside another
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("The file does not exist")
		}
	}

	fmt.Println("\n=== SENTINEL ERRORS ===")

	// Sentinel errors
	_, err = findItem("")
	if err != nil {
		if errors.Is(err, ErrInvalidInput) {
			fmt.Println("Invalid input provided")
		} else if errors.Is(err, ErrNotFound) {
			fmt.Println("Item was not found")
		} else {
			fmt.Println("Unknown error:", err)
		}
	}

	fmt.Println("\n=== ERROR HANDLING PATTERNS ===")

	// Parse integer with error handling
	num, err := parsePositiveInt("42")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed number:", num)
	}

	num, err = parsePositiveInt("-5")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed number:", num)
	}

	num, err = parsePositiveInt("abc")
	if err != nil {
		fmt.Println("Error:", err)

		// Check if specific error is wrapped
		var numErr *strconv.NumError
		if errors.As(err, &numErr) {
			fmt.Println("Failed to convert to a number:", numErr.Num)
		}
	}

	fmt.Println("\n=== PANIC AND RECOVER ===")

	// Panic and recover
	num, err = safeParse("123")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed number:", num)
	}

	num, err = safeParse("abc")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed number:", num)
	}

	fmt.Println("\n=== PRACTICAL EXAMPLES ===")

	// Demonstrate error handling in real-world scenario
	userInput := map[string]string{
		"age":      "thirty",
		"quantity": "-5",
		"email":    "invalid-email",
	}

	validateUserInput(userInput)
}

// Demonstrating error handling in a practical scenario
func validateUserInput(input map[string]string) {
	var errors []string

	// Validate age
	if ageStr, ok := input["age"]; ok {
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			errors = append(errors, fmt.Sprintf("invalid age format: %v", err))
		} else if age < 0 || age > 150 {
			errors = append(errors, fmt.Sprintf("age %d out of range", age))
		}
	}

	// Validate quantity
	if qtyStr, ok := input["quantity"]; ok {
		qty, err := strconv.Atoi(qtyStr)
		if err != nil {
			errors = append(errors, fmt.Sprintf("invalid quantity format: %v", err))
		} else if qty <= 0 {
			errors = append(errors, "quantity must be positive")
		}
	}

	// Validate email
	if email, ok := input["email"]; ok {
		if !strings.Contains(email, "@") {
			errors = append(errors, "invalid email format")
		}
	}

	// Print validation errors
	if len(errors) > 0 {
		fmt.Println("Input validation errors:")
		for i, err := range errors {
			fmt.Printf("%d. %s\n", i+1, err)
		}
	} else {
		fmt.Println("All input is valid")
	}
}

/*
Common interview questions about error handling in Go:

1. How does error handling in Go differ from exceptions in other languages?
   - Go uses explicit error return values instead of exceptions
   - Error handling is part of the normal control flow
   - No try/catch blocks, errors are checked with if statements
   - Promotes checking and handling errors at each step

2. What is the error interface in Go?
   - A built-in interface with a single method: Error() string
   - Any type that implements this method satisfies the error interface
   - Can be used to create custom error types

3. What are sentinel errors in Go?
   - Predefined error values like io.EOF or os.ErrNotExist
   - Used for specific error conditions
   - Compared with errors.Is() to check error identity
   - Gives semantic meaning to specific error cases

4. What's the difference between errors.New() and fmt.Errorf()?
   - errors.New() creates a simple error with a string message
   - fmt.Errorf() creates an error with a formatted message
   - fmt.Errorf() can also wrap errors with the %w verb (Go 1.13+)

5. What is error wrapping in Go 1.13+?
   - Allows one error to contain another
   - fmt.Errorf("... %w", err) creates a wrapped error
   - errors.Unwrap() extracts the wrapped error
   - errors.Is() and errors.As() check error chains

6. When should you use panic in Go?
   - Only for truly exceptional situations
   - When the program cannot continue at all
   - For unrecoverable errors, like failing to initialize critical components
   - Not for normal error conditions that code should handle

7. What is the purpose of recover() in Go?
   - Regain control after a panic
   - Can only be called inside a deferred function
   - Converts a panic into a normal error return
   - Used to prevent a program from crashing

8. What are best practices for error handling in Go?
   - Check errors immediately after function calls
   - Don't ignore errors unless there's a good reason
   - Add context when propagating errors
   - Use custom error types for specific error information
   - Use error wrapping to maintain error chain
   - Keep error messages clear and actionable

9. How do you check the specific type of an error?
   - Type assertion: if e, ok := err.(SomeErrorType); ok { ... }
   - Type switch: switch e := err.(type) { case SomeErrorType: ... }
   - errors.As(): if errors.As(err, &targetError) { ... }

10. What's the difference between errors.Is() and errors.As()?
    - errors.Is() checks if an error or any error it wraps matches a specific error value
    - errors.As() checks if an error or any error it wraps matches a specific error type
*/
