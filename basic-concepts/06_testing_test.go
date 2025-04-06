package main

import (
	"fmt"
	"math"
	"strings"
	"testing"
	"time"
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

// Test ValidateUser function
func TestValidateUser(t *testing.T) {
	tests := []struct {
		name        string
		user        User
		expectError bool
	}{
		{
			name: "valid user",
			user: User{
				ID:        1,
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john@example.com",
				Age:       30,
			},
			expectError: false,
		},
		{
			name: "empty first name",
			user: User{
				ID:        2,
				FirstName: "",
				LastName:  "Doe",
				Email:     "john@example.com",
				Age:       30,
			},
			expectError: true,
		},
		{
			name: "empty last name",
			user: User{
				ID:        3,
				FirstName: "John",
				LastName:  "",
				Email:     "john@example.com",
				Age:       30,
			},
			expectError: true,
		},
		{
			name: "empty email",
			user: User{
				ID:        4,
				FirstName: "John",
				LastName:  "Doe",
				Email:     "",
				Age:       30,
			},
			expectError: true,
		},
		{
			name: "negative age",
			user: User{
				ID:        5,
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john@example.com",
				Age:       -1,
			},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateUser(tc.user)

			if tc.expectError && err == nil {
				t.Errorf("ValidateUser() expected error, got nil")
			}

			if !tc.expectError && err != nil {
				t.Errorf("ValidateUser() unexpected error: %v", err)
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

// Test with mocking (simple example)
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

	// Test failure case
	failingSender := &mockEmailSender{shouldFail: true}
	err = NotifyUser(user, failingSender)
	if err == nil {
		t.Error("NotifyUser() with failing sender should return error")
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

// TestMain can be used for setup and teardown at the package level
func TestMain(m *testing.M) {
	// Setup code here
	fmt.Println("Starting tests...")

	// Run the tests
	exitCode := m.Run()

	// Teardown code here
	fmt.Println("Tests complete.")

	// Exit with the test exit code
	// In a real test, you'd use os.Exit(exitCode), but we'll omit it here
	// so it doesn't actually exit when running the tests
	// os.Exit(exitCode)
	_ = exitCode
}
