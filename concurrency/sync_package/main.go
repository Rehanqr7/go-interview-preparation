package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("=========================================")
	fmt.Println("GO SYNC PACKAGE EXAMPLES")
	fmt.Println("=========================================")

	// Mutex - mutual exclusion lock
	MutexExample()

	// RWMutex - reader/writer mutual exclusion lock
	RWMutexExample()

	// WaitGroup - wait for a collection of goroutines to finish
	WaitGroupExample()

	// Once - ensure a function is called only once
	OnceExample()

	// Atomic - atomic operations
	AtomicOperationsExample()

	// Cond - condition variable for goroutine signaling
	CondExample()

	// Map - concurrent map
	SyncMapExample()

	// Pool - object pooling
	SyncPoolExample()

	// Interview questions
	SyncPackageInterviewQuestions()
}

// Global variables (unexported to avoid conflicts)
var (
	counterVar int32
	mutexVar   sync.Mutex
	rwMutexVar sync.RWMutex
)

// MutexExample demonstrates how to use mutex for safe concurrent access
func MutexExample() {
	fmt.Println("=== MUTEX EXAMPLE ===")

	// Demonstrates race condition without mutex
	counterWithoutMutex := 0

	var wg sync.WaitGroup

	// Launch 1000 goroutines that increment the counter without mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Race condition here - no mutex protection
			counterWithoutMutex++
		}()
	}

	wg.Wait()
	fmt.Printf("Counter without mutex: %d (expected 1000)\n", counterWithoutMutex)

	// Reset counter and demonstrate with mutex
	counterWithMutex := 0

	// Launch 1000 goroutines that increment the counter with mutex protection
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Protect the counter with a mutex
			mutexVar.Lock()
			counterWithMutex++
			mutexVar.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("Counter with mutex: %d (expected 1000)\n", counterWithMutex)
	fmt.Println()
}

// RWMutexExample demonstrates how to use read-write mutex
func RWMutexExample() {
	fmt.Println("=== RWMUTEX EXAMPLE ===")

	// Data to be protected
	data := make(map[string]string)

	// Add some initial data
	rwMutexVar.Lock()
	data["key1"] = "value1"
	data["key2"] = "value2"
	data["key3"] = "value3"
	rwMutexVar.Unlock()

	var wg sync.WaitGroup

	// Launch 5 writer goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Writers need exclusive lock
			rwMutexVar.Lock()
			key := fmt.Sprintf("key%d", id+10)
			value := fmt.Sprintf("value%d", id+10)
			data[key] = value
			fmt.Printf("Writer %d: Added %s=%s\n", id, key, value)
			time.Sleep(100 * time.Millisecond) // Simulate work
			rwMutexVar.Unlock()
		}(i)
	}

	// Launch 10 reader goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Readers can share a lock
			rwMutexVar.RLock() // Read lock - multiple readers can hold this at once
			fmt.Printf("Reader %d: ", id)
			for k, v := range data {
				fmt.Printf("[%s=%s] ", k, v)
			}
			fmt.Println()
			time.Sleep(50 * time.Millisecond) // Simulate work
			rwMutexVar.RUnlock()
		}(i)
	}

	wg.Wait()
	fmt.Println()
}

// WaitGroupExample demonstrates WaitGroup for goroutine synchronization
func WaitGroupExample() {
	fmt.Println("=== WAITGROUP EXAMPLE ===")

	var wg sync.WaitGroup

	// Launch 5 goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment counter

		// Launch goroutine with id
		go func(id int) {
			defer wg.Done() // Decrement counter when done

			// Simulate work
			fmt.Printf("Worker %d starting\n", id)
			time.Sleep(time.Duration(id*200) * time.Millisecond)
			fmt.Printf("Worker %d done\n", id)
		}(i)
	}

	// Wait for all goroutines to finish
	fmt.Println("Waiting for all workers to finish...")
	wg.Wait()
	fmt.Println("All workers completed!")
	fmt.Println()
}

// OnceExample demonstrates using sync.Once for single execution
func OnceExample() {
	fmt.Println("=== ONCE EXAMPLE ===")

	// sync.Once ensures the function is called only once
	var once sync.Once

	// Initialization function to be called only once
	initFunc := func() {
		fmt.Println("Initialization function called")
		time.Sleep(200 * time.Millisecond) // Simulate work
	}

	var wg sync.WaitGroup

	// Launch 5 goroutines that all try to call the initialization
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			fmt.Printf("Goroutine %d trying to initialize...\n", id)
			once.Do(initFunc) // Only the first call will execute initFunc
			fmt.Printf("Goroutine %d done\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All initialization attempts completed")
	fmt.Println()
}

// AtomicOperationsExample demonstrates atomic operations
func AtomicOperationsExample() {
	fmt.Println("=== ATOMIC OPERATIONS EXAMPLE ===")

	// Reset the atomic counter
	atomic.StoreInt32(&counterVar, 0)

	var wg sync.WaitGroup

	// Launch 1000 goroutines to increment counter atomically
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Atomic increment
			atomic.AddInt32(&counterVar, 1)
		}()
	}

	wg.Wait()

	// Atomic load
	value := atomic.LoadInt32(&counterVar)
	fmt.Printf("Final counter value: %d\n", value)

	// Compare and swap
	oldValue := atomic.LoadInt32(&counterVar)

	// Only succeeds if counter is still oldValue
	swapped := atomic.CompareAndSwapInt32(&counterVar, oldValue, 2000)
	fmt.Printf("CAS operation success: %t, new value: %d\n", swapped, atomic.LoadInt32(&counterVar))

	// Try again with incorrect expected value
	swapped = atomic.CompareAndSwapInt32(&counterVar, oldValue, 3000)
	fmt.Printf("CAS operation success: %t, new value: %d\n", swapped, atomic.LoadInt32(&counterVar))
	fmt.Println()
}

// CondExample demonstrates condition variables
func CondExample() {
	fmt.Println("=== CONDITION VARIABLE EXAMPLE ===")

	// Create a new condition variable with its own mutex
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	// Create a queue to simulate producer-consumer pattern
	queue := make([]int, 0, 10)

	// Start consumer goroutine
	go func() {
		mu.Lock()
		defer mu.Unlock()

		// Wait until the queue has items
		for len(queue) == 0 {
			fmt.Println("Consumer: waiting for items...")
			cond.Wait() // Releases lock and waits for signal
		}

		// Consume one item
		item := queue[0]
		queue = queue[1:]
		fmt.Printf("Consumer: consumed %d\n", item)
	}()

	// Simulate some delay before producing
	time.Sleep(1 * time.Second)

	// Producer
	mu.Lock()
	fmt.Println("Producer: adding an item to queue")
	queue = append(queue, 42)
	cond.Signal() // Signal waiting consumer
	mu.Unlock()

	// Let consumer process
	time.Sleep(1 * time.Second)
	fmt.Println()
}

// SyncMapExample demonstrates the sync.Map type
func SyncMapExample() {
	fmt.Println("=== SYNC.MAP EXAMPLE ===")

	// Create a concurrent map
	var m sync.Map

	// Store values
	m.Store("key1", "value1")
	m.Store("key2", "value2")
	m.Store("key3", "value3")

	// Load a value
	value, ok := m.Load("key2")
	fmt.Printf("Load key2: value=%v, exists=%v\n", value, ok)

	// LoadOrStore (get if exists, or store and get)
	value, loaded := m.LoadOrStore("key4", "value4") // New key
	fmt.Printf("LoadOrStore key4: value=%v, previously existed=%v\n", value, loaded)

	value, loaded = m.LoadOrStore("key1", "new value") // Existing key
	fmt.Printf("LoadOrStore key1: value=%v, previously existed=%v\n", value, loaded)

	// Delete a value
	m.Delete("key2")
	value, ok = m.Load("key2")
	fmt.Printf("After delete, load key2: value=%v, exists=%v\n", value, ok)

	// Iterate over all key-value pairs
	fmt.Println("Map contents:")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("  %v: %v\n", key, value)
		return true // Continue iteration
	})
	fmt.Println()
}

// SyncPoolExample demonstrates using object pools
func SyncPoolExample() {
	fmt.Println("=== SYNC.POOL EXAMPLE ===")

	// Create a pool of byte slices
	pool := &sync.Pool{
		// New function creates a new item when Get() is called and pool is empty
		New: func() interface{} {
			buffer := make([]byte, 1024)
			fmt.Println("Creating new buffer")
			return buffer
		},
	}

	// Get a buffer from the pool (will call New)
	buffer1 := pool.Get().([]byte)
	fmt.Printf("Got buffer1 of len %d\n", len(buffer1))

	// Put the buffer back in the pool
	pool.Put(buffer1)
	fmt.Println("Put buffer1 back in pool")

	// Get a buffer again (should reuse buffer1)
	buffer2 := pool.Get().([]byte)
	fmt.Printf("Got buffer2 of len %d\n", len(buffer2))

	// Get another buffer (should call New again)
	buffer3 := pool.Get().([]byte)
	fmt.Printf("Got buffer3 of len %d\n", len(buffer3))

	// Put both buffers back
	pool.Put(buffer2)
	pool.Put(buffer3)
	fmt.Println()
}

// SyncPackageInterviewQuestions lists common interview questions about sync
func SyncPackageInterviewQuestions() {
	fmt.Println("=========================================")
	fmt.Println("COMMON INTERVIEW QUESTIONS:")
	fmt.Println("=========================================")

	fmt.Println("1. What is a mutex?")
	fmt.Println("   - Mutual exclusion lock")
	fmt.Println("   - Ensures only one goroutine accesses a resource at a time")
	fmt.Println("   - Used to protect shared data from race conditions")
	fmt.Println()

	fmt.Println("2. What is the difference between Mutex and RWMutex?")
	fmt.Println("   - Mutex: one writer at a time, blocks all readers")
	fmt.Println("   - RWMutex: allows multiple readers OR one writer")
	fmt.Println("   - RWMutex is more efficient when reads are much more frequent than writes")
	fmt.Println()

	fmt.Println("3. How does WaitGroup work?")
	fmt.Println("   - Maintains a counter of running goroutines")
	fmt.Println("   - Add(n): increment counter by n")
	fmt.Println("   - Done(): decrement counter by 1")
	fmt.Println("   - Wait(): block until counter reaches 0")
	fmt.Println()

	fmt.Println("4. What is sync.Once used for?")
	fmt.Println("   - Ensures a function is executed only once")
	fmt.Println("   - Commonly used for singleton pattern and one-time initialization")
	fmt.Println("   - Thread-safe and efficient")
	fmt.Println()

	fmt.Println("5. When would you use atomic operations vs. mutex?")
	fmt.Println("   - Atomic: simple operations on single variables (counter, flag)")
	fmt.Println("   - Mutex: complex operations or protecting multiple related variables")
	fmt.Println("   - Atomic operations are often faster but limited in scope")
	fmt.Println()

	fmt.Println("6. What is a race condition?")
	fmt.Println("   - When multiple goroutines access shared data concurrently")
	fmt.Println("   - At least one goroutine is writing")
	fmt.Println("   - The outcome depends on the timing/interleaving of operations")
	fmt.Println()

	fmt.Println("7. How can you detect race conditions in Go?")
	fmt.Println("   - Use race detector: go run -race or go test -race")
	fmt.Println("   - It detects when unsynchronized accesses to shared variables occur")
	fmt.Println()

	fmt.Println("8. What is the purpose of sync.Cond?")
	fmt.Println("   - Condition variable for goroutine signaling")
	fmt.Println("   - Used when goroutines need to wait for a condition to be true")
	fmt.Println("   - Methods: Wait, Signal, Broadcast")
	fmt.Println()

	fmt.Println("9. What is sync.Map and when would you use it?")
	fmt.Println("   - Concurrent map implementation optimized for specific access patterns")
	fmt.Println("   - Better when entries are written once but read many times")
	fmt.Println("   - Or when multiple goroutines read, write, and overwrite disjoint sets of keys")
	fmt.Println()

	fmt.Println("10. What is sync.Pool used for?")
	fmt.Println("    - Temporary object pooling/caching")
	fmt.Println("    - Reduces garbage collection pressure")
	fmt.Println("    - Useful for frequently allocated temporary objects")
	fmt.Println("    - Note: Objects may be removed from pool at any time")
	fmt.Println()

	fmt.Println("11. What is a deadlock and how can you avoid it?")
	fmt.Println("    - When goroutines are waiting for each other, forming a dependency cycle")
	fmt.Println("    - Avoid by: consistent lock ordering, timeouts, limit lock scope")
	fmt.Println("    - Go runtime detects some deadlocks and panics")
	fmt.Println()

	fmt.Println("12. What is the dining philosophers problem?")
	fmt.Println("    - Classic concurrency problem illustrating deadlock and resource contention")
	fmt.Println("    - Can be solved with mutexes, channels, or arbitrator pattern")
	fmt.Println()
}
