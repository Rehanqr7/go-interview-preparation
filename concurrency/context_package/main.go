package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	fmt.Println("=========================================")
	fmt.Println("GO CONTEXT PACKAGE EXAMPLES")
	fmt.Println("=========================================")

	// Basic context operations
	BasicContextExample()

	// Context with timeout
	DoSomethingWithTimeout()

	// Manual cancellation
	ContextWithCancellation()

	// Passing values
	ContextWithValues()

	// Deadlines
	ContextWithDeadline()

	// Propagating cancellation
	PropagatingCancellation()

	// Graceful shutdown
	GracefulShutdown()

	// HTTP requests with context
	HTTPRequestWithContext()

	// Context package overview
	ContextPackageOverview()

	// Interview questions
	ContextInterviewQuestions()
}

// BasicContextExample demonstrates creating and using a basic context
func BasicContextExample() {
	fmt.Println("=== BASIC CONTEXT EXAMPLE ===")

	// Create a background context - the root of all contexts
	ctx := context.Background()
	fmt.Printf("Background context: %v\n", ctx)

	// Create a derived context with a timeout
	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel() // Always call cancel to release resources

	// Get the deadline
	deadline, ok := timeoutCtx.Deadline()
	fmt.Printf("Context with timeout: deadline=%v, has deadline=%v\n", deadline, ok)

	// Create a derived context with a value
	valueCtx := context.WithValue(ctx, "key", "value")
	fmt.Printf("Context with value: %v\n", valueCtx)

	// Retrieve the value
	value := valueCtx.Value("key")
	fmt.Printf("Retrieved value: %v\n", value)
	fmt.Println()
}

// DoSomethingWithTimeout demonstrates using context for timeout
func DoSomethingWithTimeout() {
	fmt.Println("=== CONTEXT WITH TIMEOUT EXAMPLE ===")

	// Create a context with a 1 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Create a channel to signal completion
	done := make(chan struct{})

	// Start a goroutine that simulates a long-running operation
	go func() {
		// Simulate some work that takes 2 seconds
		fmt.Println("Starting work...")
		time.Sleep(2 * time.Second)
		fmt.Println("Work completed!") // This won't be printed due to timeout
		close(done)
	}()

	// Wait for either the work to complete or the context to timeout
	select {
	case <-done:
		fmt.Println("Work finished successfully")
	case <-ctx.Done():
		fmt.Printf("Work cancelled: %v\n", ctx.Err())
	}
	fmt.Println()
}

// ContextWithCancellation demonstrates manually cancelling a context
func ContextWithCancellation() {
	fmt.Println("=== CONTEXT WITH CANCELLATION EXAMPLE ===")

	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	// Start a worker goroutine that respects cancellation
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Worker: Received cancellation signal")
				return
			default:
				fmt.Println("Worker: Doing work...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Let the worker do some work
	time.Sleep(1500 * time.Millisecond)

	// Cancel the context
	fmt.Println("Main: Cancelling the context")
	cancel()

	// Give the worker time to respond to cancellation
	time.Sleep(1 * time.Second)
	fmt.Println()
}

// ContextWithValues demonstrates passing values through context
func ContextWithValues() {
	fmt.Println("=== CONTEXT WITH VALUES EXAMPLE ===")

	// Create a context with values
	ctx := context.Background()
	ctx = context.WithValue(ctx, "user_id", 42)
	ctx = context.WithValue(ctx, "auth_token", "secret-token")

	// Pass the context to a function
	processRequest(ctx)
	fmt.Println()
}

// processRequest is a helper function for ContextWithValues
func processRequest(ctx context.Context) {
	// Extract values from context
	userID := ctx.Value("user_id")
	token := ctx.Value("auth_token")

	fmt.Printf("Processing request for user %v with token %v\n", userID, token)

	// Pass context to another function
	validateAuth(ctx)
}

// validateAuth is a helper function for processRequest
func validateAuth(ctx context.Context) {
	// Extract token from context
	token := ctx.Value("auth_token")

	fmt.Printf("Validating authentication token: %v\n", token)
}

// ContextWithDeadline demonstrates setting a deadline
func ContextWithDeadline() {
	fmt.Println("=== CONTEXT WITH DEADLINE EXAMPLE ===")

	// Create a deadline 2 seconds from now
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// Start a worker that checks the deadline
	go func() {
		// Check if deadline has passed every 500ms
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Worker: Context done: %v\n", ctx.Err())
				return
			default:
				// Check how much time is left until deadline
				deadlineTime, ok := ctx.Deadline()
				if ok {
					timeLeft := time.Until(deadlineTime)
					fmt.Printf("Worker: Time left until deadline: %v\n", timeLeft)
				}
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Wait for the context to be done
	<-ctx.Done()
	fmt.Printf("Main: Context done: %v\n", ctx.Err())
	fmt.Println()
}

// PropagatingCancellation demonstrates propagating cancellation
func PropagatingCancellation() {
	fmt.Println("=== PROPAGATING CANCELLATION EXAMPLE ===")

	// Create a parent context
	parentCtx, parentCancel := context.WithCancel(context.Background())
	defer parentCancel()

	// Create a child context derived from parent
	childCtx, childCancel := context.WithTimeout(parentCtx, 5*time.Second)
	defer childCancel()

	// Start goroutines that use the contexts
	var wg sync.WaitGroup

	// Parent context goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-parentCtx.Done():
				fmt.Println("Parent context goroutine: cancelled")
				return
			default:
				fmt.Println("Parent context goroutine: working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Child context goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-childCtx.Done():
				fmt.Println("Child context goroutine: cancelled")
				return
			default:
				fmt.Println("Child context goroutine: working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Let goroutines run for a bit
	time.Sleep(1500 * time.Millisecond)

	// Cancel the parent context
	fmt.Println("Main: Cancelling parent context")
	parentCancel()

	// Wait for goroutines to exit
	wg.Wait()
	fmt.Println("Both goroutines exited")
	fmt.Println()
}

// GracefulShutdown demonstrates context for graceful shutdown
func GracefulShutdown() {
	fmt.Println("=== GRACEFUL SHUTDOWN EXAMPLE ===")
	fmt.Println("(This example would normally run until interrupted)")
	fmt.Println("(For demo purposes, it will auto-cancel after 3 seconds)")

	// Create a context that will be cancelled on interrupt signal
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up signal handling
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Create a channel to simulate SIGINT after 3 seconds for demo
	go func() {
		time.Sleep(3 * time.Second)
		signalChan <- syscall.SIGINT
	}()

	// Run a worker
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Worker: Shutting down gracefully")
				return
			default:
				fmt.Println("Worker: Processing...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Wait for termination signal
	sig := <-signalChan
	fmt.Printf("Received signal: %v\n", sig)

	// Cancel the context to notify workers
	fmt.Println("Main: Initiating graceful shutdown")
	cancel()

	// Give workers time to shut down
	fmt.Println("Main: Waiting for workers to finish...")
	time.Sleep(1 * time.Second)
	fmt.Println("Main: Shutdown complete")
	fmt.Println()
}

// HTTPRequestWithContext demonstrates using context with HTTP requests
func HTTPRequestWithContext() {
	fmt.Println("=== HTTP REQUEST WITH CONTEXT EXAMPLE ===")
	fmt.Println("(This example makes a real HTTP request with a timeout)")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Create a new HTTP request
	req, err := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/2", nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Execute the request
	fmt.Println("Sending HTTP request with 3 second timeout...")
	resp, err := http.DefaultClient.Do(req)

	// Check for errors
	if err != nil {
		fmt.Printf("Request error: %v\n", err)
	} else {
		defer resp.Body.Close()
		fmt.Printf("Response received: Status %s\n", resp.Status)
	}
	fmt.Println()
}

// ContextPackageOverview lists important context package elements
func ContextPackageOverview() {
	fmt.Println("=== CONTEXT PACKAGE OVERVIEW ===")

	fmt.Println("Key Functions:")
	fmt.Println("- context.Background(): Root context, never cancelled")
	fmt.Println("- context.TODO(): Placeholder when context not available")
	fmt.Println("- WithCancel(): Returns cancellable context and cancel function")
	fmt.Println("- WithDeadline(): Context that cancels at specified time")
	fmt.Println("- WithTimeout(): Context that cancels after duration")
	fmt.Println("- WithValue(): Context with key-value data")

	fmt.Println("\nContext Interface Methods:")
	fmt.Println("- Deadline(): Returns deadline and if set")
	fmt.Println("- Done(): Returns channel that's closed when cancelled")
	fmt.Println("- Err(): Returns error explaining why Done closed")
	fmt.Println("- Value(): Returns value for key")

	fmt.Println("\nBest Practices:")
	fmt.Println("- Always call cancel() to release resources")
	fmt.Println("- Pass context as first parameter")
	fmt.Println("- Don't store contexts in structs")
	fmt.Println("- Use WithValue only for request-scoped data")
	fmt.Println("- Keep key types private to packages")
	fmt.Println()
}

// ContextInterviewQuestions lists common interview questions about context
func ContextInterviewQuestions() {
	fmt.Println("=========================================")
	fmt.Println("COMMON INTERVIEW QUESTIONS:")
	fmt.Println("=========================================")

	fmt.Println("1. What is the context package and why is it used?")
	fmt.Println("   - Package for propagating deadlines, cancellation signals, and request values")
	fmt.Println("   - Used to control timeouts, cancellation, and carry request-scoped values")
	fmt.Println("   - Helps prevent resource leaks and implement graceful shutdown")
	fmt.Println()

	fmt.Println("2. What are the two root context types and when to use each?")
	fmt.Println("   - context.Background(): Root of all contexts, used in main/init/tests")
	fmt.Println("   - context.TODO(): Placeholder when it's unclear which context to use")
	fmt.Println()

	fmt.Println("3. How does context cancellation propagate?")
	fmt.Println("   - When a context is cancelled, all contexts derived from it are cancelled")
	fmt.Println("   - Allows for cancelling entire subtrees of operations")
	fmt.Println("   - Child contexts can't affect parent contexts")
	fmt.Println()

	fmt.Println("4. How do you handle timeouts with context?")
	fmt.Println("   - Use WithTimeout or WithDeadline to create a context with time constraints")
	fmt.Println("   - Operations using this context can check ctx.Done() for timeout")
	fmt.Println("   - Common in HTTP servers, DB operations, and API calls")
	fmt.Println()

	fmt.Println("5. What are best practices for passing values in context?")
	fmt.Println("   - Only use for request-scoped data (tracing ID, auth tokens)")
	fmt.Println("   - Don't use for passing optional parameters")
	fmt.Println("   - Use custom key types (not strings) to avoid collisions")
	fmt.Println("   - Keep keys as unexported types")
	fmt.Println()

	fmt.Println("6. How would you implement a function that respects cancellation?")
	fmt.Println("   - Accept context as first parameter")
	fmt.Println("   - Check ctx.Done() in loops or long operations")
	fmt.Println("   - Return quickly when context is cancelled")
	fmt.Println("   - Return ctx.Err() or wrap it in custom error")
	fmt.Println()

	fmt.Println("7. What's the relationship between context and http.Request?")
	fmt.Println("   - http.Request has a Context() method that returns its context")
	fmt.Println("   - Context automatically cancelled when handler returns")
	fmt.Println("   - Use req.WithContext() to create new request with modified context")
	fmt.Println()

	fmt.Println("8. What are common mistakes with context?")
	fmt.Println("   - Storing context in structs")
	fmt.Println("   - Not calling cancel() function")
	fmt.Println("   - Using context.Value for function parameters")
	fmt.Println("   - Creating many child contexts instead of siblings")
	fmt.Println()

	fmt.Println("9. How do you implement graceful shutdown with context?")
	fmt.Println("   - Trap termination signals (SIGINT/SIGTERM)")
	fmt.Println("   - Cancel a context when signal received")
	fmt.Println("   - Pass this context to subsystems")
	fmt.Println("   - Each component checks ctx.Done() and shuts down when triggered")
	fmt.Println()

	fmt.Println("10. How do you unit test code that uses context?")
	fmt.Println("    - Test timeout by using a short WithTimeout context")
	fmt.Println("    - Test cancellation by creating context and calling cancel()")
	fmt.Println("    - Test values by using WithValue and checking behavior")
	fmt.Println("    - Use context.Background() for tests without deadline constraints")
	fmt.Println()
}
