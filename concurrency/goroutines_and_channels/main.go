package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("=========================================")
	fmt.Println("GO GOROUTINES AND CHANNELS EXAMPLES")
	fmt.Println("=========================================")

	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Basic goroutine
	SimpleGoroutine()

	// WaitGroup for synchronization
	WaitGroupExample()

	// Channel examples
	UnbufferedChannels()
	BufferedChannels()
	ChannelDirections()
	ClosingChannels()
	IteratingOverChannels()

	// Select examples
	SelectStatement()
	SelectWithTimeout()
	SelectWithDefault()

	// Concurrency patterns
	WorkerPool()
	FanOutFanIn()

	// Informational
	ChannelComparison()

	// Interview questions
	GoroutinesAndChannelsInterviewQuestions()
}

// SimpleGoroutine demonstrates a basic goroutine
func SimpleGoroutine() {
	fmt.Println("=== SIMPLE GOROUTINE EXAMPLE ===")

	// Start a goroutine
	go func() {
		fmt.Println("Hello from goroutine!")
	}()

	// Give the goroutine time to execute
	// In real code, you'd use proper synchronization instead
	time.Sleep(100 * time.Millisecond)
	fmt.Println()
}

// WaitGroupExample demonstrates using WaitGroup for synchronization
func WaitGroupExample() {
	fmt.Println("=== WAITGROUP EXAMPLE ===")

	var wg sync.WaitGroup

	// Launch 5 workers
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment counter before launching goroutine

		go func(id int) {
			defer wg.Done() // Decrement counter when goroutine completes

			fmt.Printf("Worker %d starting\n", id)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("Worker %d done\n", id)
		}(i)
	}

	fmt.Println("Waiting for all workers to complete...")
	wg.Wait() // Block until counter becomes 0
	fmt.Println("All workers completed!")
	fmt.Println()
}

// UnbufferedChannels demonstrates basic channel operations
func UnbufferedChannels() {
	fmt.Println("=== UNBUFFERED CHANNELS EXAMPLE ===")

	// Create an unbuffered channel
	ch := make(chan string)

	// Sender goroutine
	go func() {
		fmt.Println("Sender: Sending message")
		ch <- "Hello from sender!" // Will block until someone receives
		fmt.Println("Sender: Message sent")
	}()

	// Give sender time to start
	time.Sleep(100 * time.Millisecond)

	// Receive the message
	fmt.Println("Receiver: About to receive")
	msg := <-ch // Will unblock the sender
	fmt.Printf("Receiver: Got message: %q\n", msg)
	fmt.Println()
}

// BufferedChannels demonstrates buffered channel behavior
func BufferedChannels() {
	fmt.Println("=== BUFFERED CHANNELS EXAMPLE ===")

	// Create a buffered channel with capacity 2
	ch := make(chan string, 2)

	// Send messages (won't block until buffer is full)
	fmt.Println("Sending to buffered channel")
	ch <- "First message"
	fmt.Println("Sent first message")
	ch <- "Second message"
	fmt.Println("Sent second message")

	// This would block because buffer is full:
	// ch <- "Third message"

	// Receive messages
	fmt.Printf("Received: %q\n", <-ch)
	fmt.Printf("Received: %q\n", <-ch)
	fmt.Println()
}

// ChannelDirections demonstrates channel direction constraints
func ChannelDirections() {
	fmt.Println("=== CHANNEL DIRECTIONS EXAMPLE ===")

	// Create a bidirectional channel
	ch := make(chan int)

	// Start the sender
	go sender(ch)

	// Start the receiver, passing the same channel
	go receiver(ch)

	// Let them communicate
	time.Sleep(1 * time.Second)
	fmt.Println()
}

// sender only sends to the channel (send-only channel)
func sender(ch chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Sender: sending %d\n", i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
}

// receiver only receives from the channel (receive-only channel)
func receiver(ch <-chan int) {
	for i := 0; i < 5; i++ {
		val := <-ch
		fmt.Printf("Receiver: got %d\n", val)
	}
}

// ClosingChannels demonstrates closing a channel and ranging over it
func ClosingChannels() {
	fmt.Println("=== CLOSING CHANNELS EXAMPLE ===")

	// Create a channel
	ch := make(chan int, 5)

	// Launch producer
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Printf("Sent: %d\n", i)
		}
		fmt.Println("Producer: closing channel")
		close(ch)
	}()

	// Consumer: loop until channel is closed
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel closed, exiting")
			break
		}
		fmt.Printf("Received: %d\n", val)
	}
	fmt.Println()
}

// IteratingOverChannels demonstrates the range syntax for channels
func IteratingOverChannels() {
	fmt.Println("=== ITERATING OVER CHANNELS EXAMPLE ===")

	// Create a channel
	ch := make(chan int, 5)

	// Launch producer
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // Important: close channel when done sending
	}()

	// Consume with range (automatically handles closed channels)
	for val := range ch {
		fmt.Printf("Received: %d\n", val)
	}
	fmt.Println("Channel closed, loop exited")
	fmt.Println()
}

// SelectStatement demonstrates using select to wait on multiple channels
func SelectStatement() {
	fmt.Println("=== SELECT STATEMENT EXAMPLE ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// First goroutine sends to ch1 after 100ms
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Message from channel 1"
	}()

	// Second goroutine sends to ch2 after 50ms
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "Message from channel 2"
	}()

	// Use select to wait for messages from either channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received from ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from ch2: %s\n", msg2)
		}
	}
	fmt.Println()
}

// SelectWithTimeout demonstrates using select with timeout
func SelectWithTimeout() {
	fmt.Println("=== SELECT WITH TIMEOUT EXAMPLE ===")

	ch := make(chan string)

	// Goroutine that will send a message after 500ms
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- "Message from goroutine"
	}()

	// Wait for the message with a 250ms timeout
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(250 * time.Millisecond):
		fmt.Println("Timeout! No message received in time")
	}

	// Wait for the message with a 1 second timeout (will succeed)
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout! No message received in time")
	}
	fmt.Println()
}

// SelectWithDefault demonstrates non-blocking channel operations
func SelectWithDefault() {
	fmt.Println("=== SELECT WITH DEFAULT EXAMPLE ===")

	ch := make(chan string)

	// Try to receive, but don't block
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("No message available")
	}

	// Try to send, but don't block
	select {
	case ch <- "Hello":
		fmt.Println("Sent message")
	default:
		fmt.Println("No receiver available")
	}
	fmt.Println()
}

// WorkerPool demonstrates a worker pool pattern
func WorkerPool() {
	fmt.Println("=== WORKER POOL EXAMPLE ===")

	const numJobs = 10
	const numWorkers = 3

	// Create job and result channels
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, results)
		}(w)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // No more jobs

	// Start a goroutine to close results channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
	fmt.Println()
}

// worker processes jobs from jobs channel and sends results to results channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		results <- job * 2 // Simulate some processing
	}
}

// FanOutFanIn demonstrates the fan-out/fan-in pattern
func FanOutFanIn() {
	fmt.Println("=== FAN-OUT/FAN-IN EXAMPLE ===")

	// Create channels
	input := make(chan int, 10)

	// Send input values
	go func() {
		for i := 0; i < 10; i++ {
			input <- i
		}
		close(input)
	}()

	// Create multiple channels to fan out the work
	c1 := fanOut(input)
	c2 := fanOut(input)
	c3 := fanOut(input)

	// Fan in the results
	for result := range fanIn(c1, c2, c3) {
		fmt.Printf("Result: %d\n", result)
	}
	fmt.Println()
}

// fanOut creates a channel that processes input values and sends results
func fanOut(input <-chan int) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)
		for n := range input {
			// Simulate varying processing times
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			output <- n * n // Square the number
		}
	}()

	return output
}

// fanIn multiplexes multiple input channels onto a single output channel
func fanIn(inputs ...<-chan int) <-chan int {
	output := make(chan int)
	var wg sync.WaitGroup

	// For each input channel, start a goroutine to forward values
	for _, ch := range inputs {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				output <- n
			}
		}(ch)
	}

	// When all inputs are drained, close the output
	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

// ChannelComparison demonstrates channel behaviors and differences
func ChannelComparison() {
	fmt.Println("=== CHANNEL COMPARISON ===")

	fmt.Println("Channel Behaviors:")
	fmt.Println("- Unbuffered: send blocks until receive")
	fmt.Println("- Buffered: send blocks only when buffer full")
	fmt.Println("- Receive always blocks if no data available")
	fmt.Println("- Channel closing: sends panic, receives get zero value")
	fmt.Println("- Closed check: val, ok := <-ch (ok is false if closed)")
	fmt.Println("- nil channels: sends and receives block forever")

	fmt.Println("\nChannel Use Cases:")
	fmt.Println("- Signaling: close a channel to broadcast to multiple goroutines")
	fmt.Println("- Done channel: channel to signal completion")
	fmt.Println("- Worker pools: distribute work among multiple workers")
	fmt.Println("- Rate limiting: buffer capacity controls processing rate")
	fmt.Println("- Pipelines: chain of stages connected by channels")
	fmt.Println()
}

// GoroutinesAndChannelsInterviewQuestions lists common interview questions
func GoroutinesAndChannelsInterviewQuestions() {
	fmt.Println("=========================================")
	fmt.Println("COMMON INTERVIEW QUESTIONS:")
	fmt.Println("=========================================")

	fmt.Println("1. What is a goroutine and how is it different from a thread?")
	fmt.Println("   - Lightweight thread managed by Go runtime")
	fmt.Println("   - Much smaller stack size (2KB initially vs MB for OS threads)")
	fmt.Println("   - Cheaper creation and context switching")
	fmt.Println("   - Go runtime multiplexes goroutines onto OS threads")
	fmt.Println()

	fmt.Println("2. How do goroutines communicate?")
	fmt.Println("   - Primarily through channels")
	fmt.Println("   - Can also use shared memory with proper synchronization")
	fmt.Println("   - \"Don't communicate by sharing memory; share memory by communicating\"")
	fmt.Println()

	fmt.Println("3. What is the difference between buffered and unbuffered channels?")
	fmt.Println("   - Unbuffered: synchronous, sender blocks until receiver receives")
	fmt.Println("   - Buffered: asynchronous up to buffer capacity")
	fmt.Println("   - Buffered channels decouple sender and receiver temporally")
	fmt.Println()

	fmt.Println("4. How do you prevent goroutine leaks?")
	fmt.Println("   - Ensure goroutines can exit (e.g., by using contexts, timeout channels)")
	fmt.Println("   - Properly close channels to signal completion")
	fmt.Println("   - Use cancellation mechanisms like context.Context")
	fmt.Println("   - Use WaitGroups to track completion")
	fmt.Println()

	fmt.Println("5. What does the select statement do?")
	fmt.Println("   - Waits on multiple channel operations")
	fmt.Println("   - Blocks until one case can proceed")
	fmt.Println("   - If multiple cases ready, chooses one at random")
	fmt.Println("   - default case makes select non-blocking")
	fmt.Println()

	fmt.Println("6. What happens when you close a channel?")
	fmt.Println("   - Sends on closed channel panic")
	fmt.Println("   - Receives from closed channel get zero value immediately")
	fmt.Println("   - Receive check (val, ok := <-ch) returns ok=false when closed")
	fmt.Println()

	fmt.Println("7. What is a race condition and how to detect it?")
	fmt.Println("   - Concurrent access to shared data without proper synchronization")
	fmt.Println("   - Detect with go run -race or go test -race")
	fmt.Println("   - Fix with proper synchronization mechanisms")
	fmt.Println()

	fmt.Println("8. What are common concurrency patterns in Go?")
	fmt.Println("   - Worker pools: fixed number of workers processing from queue")
	fmt.Println("   - Fan-out/fan-in: distribute work and collect results")
	fmt.Println("   - Pipeline: chain of stages connected by channels")
	fmt.Println("   - Cancellation: propagate cancellation using context")
	fmt.Println()

	fmt.Println("9. How many OS threads does Go use for goroutines?")
	fmt.Println("   - By default, GOMAXPROCS (usually matches CPU cores)")
	fmt.Println("   - Can be modified with runtime.GOMAXPROCS()")
	fmt.Println("   - Current value is", runtime.GOMAXPROCS(0))
	fmt.Println()

	fmt.Println("10. What is channel directionality?")
	fmt.Println("    - Restrict channel to send-only (chan<-) or receive-only (<-chan)")
	fmt.Println("    - Provides better type safety")
	fmt.Println("    - Documents intent and prevents incorrect usage")
	fmt.Println()

	fmt.Println("11. What's the difference between len() and cap() for channels?")
	fmt.Println("    - len(): number of elements currently in the channel")
	fmt.Println("    - cap(): total capacity of the channel buffer")
	fmt.Println()

	fmt.Println("12. When would you use a nil channel?")
	fmt.Println("    - In select statements to disable specific cases")
	fmt.Println("    - Note: sends and receives on nil channels block forever")
	fmt.Println()
}
