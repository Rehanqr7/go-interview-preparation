package main

import (
	"fmt"
	"math"
)

// Functions to be tested
// These functions would typically be in a separate package

// Sum returns the sum of two integers
func Sum(a, b int) int {
	return a + b
}

// Multiply returns the product of two integers
func Multiply(a, b int) int {
	return a * b
}

// CircleArea returns the area of a circle with the given radius
func CircleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("negative radius: %f", radius)
	}
	return math.Pi * radius * radius, nil
}

// WordCount counts the number of words in a string
func WordCount(s string) int {
	// Edge case: empty string
	if len(s) == 0 {
		return 0
	}

	// Count spaces to determine words
	count := 1 // Start with 1 for the first word
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && i > 0 && s[i-1] != ' ' {
			count++
		}
	}

	// Handle cases like "  hello  "
	if s[0] == ' ' {
		count--
	}
	if s[len(s)-1] == ' ' {
		count--
	}

	return count
}

// User represents a user in the system
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Age       int
}

// ValidateUser checks if user data is valid
func ValidateUser(u User) error {
	if u.FirstName == "" {
		return fmt.Errorf("first name cannot be empty")
	}
	if u.LastName == "" {
		return fmt.Errorf("last name cannot be empty")
	}
	if u.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	if u.Age < 0 {
		return fmt.Errorf("age cannot be negative")
	}
	return nil
}

// EmailSender is an interface for sending emails
type EmailSender interface {
	Send(email, subject, body string) error
}

// NotifyUser sends a notification email to a user
func NotifyUser(user User, sender EmailSender) error {
	body := fmt.Sprintf("Hello %s, your account has been created.", user.FirstName)
	return sender.Send(user.Email, "Account Created", body)
}

func main() {
	fmt.Println("=== TESTING IN GO ===")

	fmt.Println("To run tests in this package, use:")
	fmt.Println("    go test -v")
	fmt.Println("To run specific tests:")
	fmt.Println("    go test -v -run TestSum")
	fmt.Println("To run benchmarks:")
	fmt.Println("    go test -v -bench=.")
	fmt.Println("To run examples:")
	fmt.Println("    go test -v -run Example")
	fmt.Println("To see test coverage:")
	fmt.Println("    go test -cover")
	fmt.Println("    go test -coverprofile=coverage.out")
	fmt.Println("    go tool cover -html=coverage.out")

	// Demonstrate the functions
	fmt.Println("\nDemonstrating functions that would be tested:")

	fmt.Printf("Sum(5, 3) = %d\n", Sum(5, 3))
	fmt.Printf("Multiply(4, 6) = %d\n", Multiply(4, 6))

	area, err := CircleArea(5.0)
	if err != nil {
		fmt.Printf("Error calculating circle area: %v\n", err)
	} else {
		fmt.Printf("CircleArea(5.0) = %.2f\n", area)
	}

	fmt.Printf("WordCount(\"hello world\") = %d\n", WordCount("hello world"))

	// Create a user and validate
	user := User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Age:       30,
	}

	err = ValidateUser(user)
	if err != nil {
		fmt.Printf("User validation error: %v\n", err)
	} else {
		fmt.Println("User is valid!")
	}

	// Create invalid user
	invalidUser := User{
		ID:    2,
		Email: "jane@example.com",
		Age:   25,
	}

	err = ValidateUser(invalidUser)
	if err != nil {
		fmt.Printf("Invalid user validation error: %v\n", err)
	} else {
		fmt.Println("User is valid!")
	}
}

// TESTS FOR THESE FUNCTIONS WOULD TYPICALLY BE IN A FILE NAMED *_test.go
// For example, the tests for this file would be in testing_test.go:

/*
// Example of testing_test.go file:

package main

import (
	"testing"
	"math"
)

// Basic unit test
func TestSum(t *testing.T) {
	// Table-driven test
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed numbers", -2, 3, 1},
		{"zeros", 0, 0, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Sum(tc.a, tc.b)
			if got != tc.expected {
				t.Errorf("Sum(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}

// Test with error checking
func TestCircleArea(t *testing.T) {
	// Test cases
	tests := []struct {
		name        string
		radius      float64
		expected    float64
		expectError bool
	}{
		{"positive radius", 5.0, math.Pi * 25.0, false},
		{"zero radius", 0.0, 0.0, false},
		{"negative radius", -5.0, 0.0, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CircleArea(tc.radius)

			// Check error expectations
			if tc.expectError && err == nil {
				t.Errorf("CircleArea(%f) expected error, got nil", tc.radius)
				return
			}
			if !tc.expectError && err != nil {
				t.Errorf("CircleArea(%f) unexpected error: %v", tc.radius, err)
				return
			}

			// If we don't expect an error, check the result
			if !tc.expectError {
				if math.Abs(got-tc.expected) > 1e-10 {
					t.Errorf("CircleArea(%f) = %f; want %f", tc.radius, got, tc.expected)
				}
			}
		})
	}
}

// Testing with setup and teardown
func TestWithSetupAndTeardown(t *testing.T) {
	// Setup
	t.Log("Setting up test environment")
	// Could initialize a database, create files, etc.

	// Cleanup with defer
	defer func() {
		t.Log("Tearing down test environment")
		// Could close database connections, delete files, etc.
	}()

	// Actual test
	t.Run("test case", func(t *testing.T) {
		// Test logic here
		if 1+1 != 2 {
			t.Error("Something is seriously wrong with math")
		}
	})
}

// Testing with subtests
func TestWordCount(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello world", 2},
		{"", 0},
		{"oneword", 1},
		{"   spaced   words   ", 2},
		{"1 2 3 4 5", 5},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("case %d: %q", i, tc.input), func(t *testing.T) {
			got := WordCount(tc.input)
			if got != tc.expected {
				t.Errorf("WordCount(%q) = %d; want %d", tc.input, got, tc.expected)
			}
		})
	}
}

// Benchmark example
func BenchmarkSum(b *testing.B) {
	// Run the Sum function b.N times
	for i := 0; i < b.N; i++ {
		Sum(4, 5)
	}
}

// Benchmark with setup
func BenchmarkWordCount(b *testing.B) {
	longText := strings.Repeat("The quick brown fox jumps over the lazy dog. ", 100)

	// Reset the timer to exclude setup time
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		WordCount(longText)
	}
}

// Example function that will be verified by running the test
func ExampleSum() {
	result := Sum(1, 2)
	fmt.Println(result)
	// Output: 3
}

// Another example with multiple lines of output
func ExampleWordCount() {
	count1 := WordCount("hello world")
	fmt.Println(count1)

	count2 := WordCount("one two three")
	fmt.Println(count2)

	// Output:
	// 2
	// 3
}

// Mock implementation for email sender
type mockEmailSender struct {
	sentEmails []struct {
		email   string
		subject string
		body    string
	}
	shouldFail bool
}

func (m *mockEmailSender) Send(email, subject, body string) error {
	if m.shouldFail {
		return fmt.Errorf("failed to send email")
	}
	m.sentEmails = append(m.sentEmails, struct {
		email   string
		subject string
		body    string
	}{email, subject, body})
	return nil
}

func TestNotifyUser(t *testing.T) {
	// Create a mock email sender
	mockSender := &mockEmailSender{}

	// Create a test user
	user := User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Age:       30,
	}

	// Call the function with our mock
	err := NotifyUser(user, mockSender)

	// Verify no error occurred
	if err != nil {
		t.Errorf("NotifyUser() returned error: %v", err)
	}

	// Verify an email was sent
	if len(mockSender.sentEmails) != 1 {
		t.Errorf("Expected 1 email to be sent, got %d", len(mockSender.sentEmails))
	}

	// Verify the email content
	if len(mockSender.sentEmails) > 0 {
		sent := mockSender.sentEmails[0]
		if sent.email != user.Email {
			t.Errorf("Wrong recipient email: got %s, want %s", sent.email, user.Email)
		}
		if sent.subject != "Account Created" {
			t.Errorf("Wrong subject: got %s, want %s", sent.subject, "Account Created")
		}
		expectedBody := fmt.Sprintf("Hello %s, your account has been created.", user.FirstName)
		if sent.body != expectedBody {
			t.Errorf("Wrong body: got %s, want %s", sent.body, expectedBody)
		}
	}
}

// Parallel tests
func TestParallel(t *testing.T) {
	// Tests in this group will run in parallel with each other
	t.Run("group", func(t *testing.T) {
		t.Run("first", func(t *testing.T) {
			t.Parallel()
			time.Sleep(100 * time.Millisecond)
		})
		t.Run("second", func(t *testing.T) {
			t.Parallel()
			time.Sleep(100 * time.Millisecond)
		})
		t.Run("third", func(t *testing.T) {
			t.Parallel()
			time.Sleep(100 * time.Millisecond)
		})
	})
}
*/

/*
Common interview questions about testing in Go:

1. How do you write a basic unit test in Go?
   - Create a file ending with _test.go in the same package
   - Write functions in the format TestXxx(t *testing.T)
   - Run with "go test"

2. What are table-driven tests in Go?
   - Tests that define a slice/array of test cases
   - Each case has inputs and expected outputs
   - Tests run in a loop over all test cases
   - Very common pattern in Go testing

3. How do you run a specific test in Go?
   - Use the -run flag: go test -run TestXxx
   - Can use regex patterns to match multiple tests

4. What's the difference between tests, benchmarks, and examples?
   - Tests (TestXxx): Verify functionality
   - Benchmarks (BenchmarkXxx): Measure performance
   - Examples (ExampleXxx): Document usage and test correctness

5. How do you measure test coverage in Go?
   - Run: go test -cover
   - For detailed output: go test -coverprofile=coverage.out
   - Visualize: go tool cover -html=coverage.out

6. How do you implement setup and teardown in Go tests?
   - Use test fixtures/helpers for setup
   - Use defer for teardown
   - TestMain(m *testing.M) for package-level setup/teardown

7. What are subtests in Go?
   - Tests created with t.Run("name", func(t *testing.T) {...})
   - Allow grouping related tests
   - Can run specific subtests with -run TestName/SubtestName

8. How do you benchmark in Go?
   - Create functions in the format BenchmarkXxx(b *testing.B)
   - Run the code b.N times in a loop
   - Run with "go test -bench=."

9. How do you write examples in Go?
   - Create functions in the format ExampleXxx()
   - Include expected output in a comment: // Output: expected
   - Serve as both tests and documentation

10. How do you mock dependencies in Go tests?
    - Use interfaces for external dependencies
    - Create mock implementations for testing
    - Inject the mock during tests

11. How do parallel tests work in Go?
    - Call t.Parallel() in a test to indicate it can run in parallel
    - Use -parallel N flag to specify max parallelism
    - Helps test concurrent code and reduces test execution time

12. What test frameworks/libraries are commonly used in Go?
    - Built-in testing package is most common
    - testify for assertions and mocking
    - gomock for mocking
    - httptest for HTTP testing
    - go-sqlmock for database mocking
*/
