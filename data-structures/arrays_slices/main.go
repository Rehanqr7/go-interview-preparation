package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("=========================================")
	fmt.Println("GO ARRAYS AND SLICES EXAMPLES")
	fmt.Println("=========================================")

	// Basic arrays
	BasicArraysExample()

	// Basic slices
	BasicSlicesExample()

	// Slice manipulation
	SliceManipulationExample()

	// Slice capacity and growth
	SliceCapacityExample()

	// Slice memory sharing
	SliceMemorySharingExample()

	// Multidimensional slices
	MultidimensionalSlicesExample()

	// Slice sorting
	SliceSortingExample()

	// Common slice operations
	CommonSliceOperationsExample()

	// Performance considerations
	PerformanceConsiderationsExample()

	// Interview questions
	ArraysAndSlicesInterviewQuestions()
}

// BasicArraysExample demonstrates array declaration and use
func BasicArraysExample() {
	fmt.Println("=== BASIC ARRAYS EXAMPLE ===")

	// Fixed-size array declaration and initialization
	var numbers [5]int                                    // Zero-initialized [0,0,0,0,0]
	primes := [5]int{2, 3, 5, 7, 11}                      // Initialize with values
	fibonacci := [...]int{1, 1, 2, 3, 5, 8, 13}           // Size determined by initializer
	colors := [4]string{"red", "blue", "green", "yellow"} // String array

	fmt.Println("Empty array:", numbers)
	fmt.Println("Prime numbers:", primes)
	fmt.Println("Fibonacci numbers:", fibonacci)
	fmt.Println("Colors:", colors)

	// Accessing array elements
	fmt.Println("First prime:", primes[0])
	fmt.Println("Last color:", colors[len(colors)-1])

	// Modifying elements
	numbers[0] = 10
	numbers[1] = 20
	fmt.Println("Modified numbers:", numbers)

	// Array length
	fmt.Println("Length of primes array:", len(primes))
	fmt.Println("Length of fibonacci array:", len(fibonacci))

	// Iterating over an array
	fmt.Println("Iterating over colors array:")
	for i, color := range colors {
		fmt.Printf("  Index %d: %s\n", i, color)
	}

	// Comparing arrays
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{3, 2, 1}
	fmt.Println("arr1 == arr2:", arr1 == arr2) // true - same values
	fmt.Println("arr1 == arr3:", arr1 == arr3) // false - different values

	fmt.Println()
}

// BasicSlicesExample demonstrates slice declaration and use
func BasicSlicesExample() {
	fmt.Println("=== BASIC SLICES EXAMPLE ===")

	// Slice declaration and initialization
	var emptySlice []int                         // nil slice
	numbers := []int{1, 2, 3, 4, 5}              // Initialized slice
	names := []string{"Alice", "Bob", "Charlie"} // String slice

	fmt.Println("Empty slice:", emptySlice, "- Length:", len(emptySlice), "- Is nil:", emptySlice == nil)
	fmt.Println("Numbers slice:", numbers, "- Length:", len(numbers))
	fmt.Println("Names slice:", names, "- Length:", len(names))

	// Creating a slice from an array
	primes := [6]int{2, 3, 5, 7, 11, 13}
	somePrimes := primes[1:4]      // Creates slice [3,5,7]
	allPrimes := primes[:]         // Slice of whole array
	firstThreePrimes := primes[:3] // Slice [2,3,5]
	lastThreePrimes := primes[3:]  // Slice [7,11,13]

	fmt.Println("Original array:", primes)
	fmt.Println("Slice somePrimes:", somePrimes)
	fmt.Println("Slice allPrimes:", allPrimes)
	fmt.Println("Slice firstThreePrimes:", firstThreePrimes)
	fmt.Println("Slice lastThreePrimes:", lastThreePrimes)

	// Creating a slice with make
	sliceWithMake := make([]int, 5) // Length 5, capacity 5
	fmt.Println("Slice created with make:", sliceWithMake)

	sliceWithCapacity := make([]int, 3, 10) // Length 3, capacity 10
	fmt.Println("Slice with capacity:", sliceWithCapacity,
		"- Length:", len(sliceWithCapacity),
		"- Capacity:", cap(sliceWithCapacity))

	fmt.Println()
}

// SliceManipulationExample demonstrates common slice operations
func SliceManipulationExample() {
	fmt.Println("=== SLICE MANIPULATION EXAMPLE ===")

	// Appending elements
	numbers := []int{1, 2, 3}
	fmt.Println("Original slice:", numbers)

	numbers = append(numbers, 4)
	fmt.Println("After append(numbers, 4):", numbers)

	numbers = append(numbers, 5, 6, 7)
	fmt.Println("After append(numbers, 5, 6, 7):", numbers)

	// Appending another slice
	moreNumbers := []int{8, 9, 10}
	numbers = append(numbers, moreNumbers...) // The ellipsis ... is required
	fmt.Println("After appending another slice:", numbers)

	// Removing elements
	// From the start (shift)
	numbers = numbers[1:]
	fmt.Println("After removing first element:", numbers)

	// From the end (pop)
	numbers = numbers[:len(numbers)-1]
	fmt.Println("After removing last element:", numbers)

	// From the middle (removing element at index 2)
	i := 2
	numbers = append(numbers[:i], numbers[i+1:]...)
	fmt.Println("After removing element at index 2:", numbers)

	// Clearing a slice
	numbers = numbers[:0] // Keep the same capacity
	fmt.Println("After clearing the slice:", numbers, "- Length:", len(numbers), "- Capacity:", cap(numbers))

	fmt.Println()
}

// SliceCapacityExample demonstrates slice growth and capacity
func SliceCapacityExample() {
	fmt.Println("=== SLICE CAPACITY AND GROWTH EXAMPLE ===")

	// Creating a slice with a specific capacity
	s := make([]int, 0, 5)
	fmt.Printf("Initial slice: len=%d cap=%d %v\n", len(s), cap(s), s)

	// Appending elements and observing capacity changes
	for i := 1; i <= 10; i++ {
		s = append(s, i)
		fmt.Printf("After appending %d: len=%d cap=%d %v\n", i, len(s), cap(s), s)
	}

	// Go doubles capacity when the slice needs to grow
	// This is an amortized O(1) operation

	// Pre-allocating for better performance
	s = make([]int, 0, 1000)
	fmt.Printf("Pre-allocated slice: len=%d cap=%d\n", len(s), cap(s))

	// This append won't reallocate
	for i := 0; i < 1000; i++ {
		s = append(s, i)
	}
	fmt.Printf("After 1000 appends: len=%d cap=%d\n", len(s), cap(s))

	fmt.Println()
}

// SliceMemorySharingExample demonstrates how slices share memory
func SliceMemorySharingExample() {
	fmt.Println("=== SLICE MEMORY SHARING EXAMPLE ===")

	// Original slice
	original := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", original)

	// Creating a slice from another slice
	shared := original[1:4]
	fmt.Println("Shared slice:", shared)

	// Modifying the shared slice affects the original
	shared[0] = 99
	fmt.Println("Original after modifying shared:", original)
	fmt.Println("Shared after modification:", shared)

	// When appending to a shared slice, it may detach from the original
	// if it exceeds the capacity
	fmt.Println("Shared capacity:", cap(shared))
	shared = append(shared, 100)

	// shared may or may not affect original here depending on capacity
	fmt.Println("Original after append to shared:", original)
	fmt.Println("Shared after append:", shared)

	// Making a true copy
	numbers := []int{1, 2, 3, 4, 5}
	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, numbers)

	// Modifying the copy doesn't affect the original
	numbersCopy[0] = 99
	fmt.Println("Original numbers:", numbers)
	fmt.Println("Copy after modification:", numbersCopy)

	fmt.Println()
}

// MultidimensionalSlicesExample demonstrates multi-dimensional slices
func MultidimensionalSlicesExample() {
	fmt.Println("=== MULTIDIMENSIONAL SLICES EXAMPLE ===")

	// 2D slice (slice of slices)
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("2D matrix:", matrix)

	// Accessing elements
	fmt.Println("Element at row 1, col 2:", matrix[1][2])

	// Modifying elements
	matrix[0][0] = 99
	fmt.Println("Modified matrix:", matrix)

	// Creating a 2D slice with make
	rows, cols := 3, 4
	grid := make([][]int, rows)

	for i := 0; i < rows; i++ {
		grid[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			grid[i][j] = i*cols + j
		}
	}
	fmt.Println("Created grid:", grid)

	// Iterating over a 2D slice
	fmt.Println("Iterating over the grid:")
	for i, row := range grid {
		for j, val := range row {
			fmt.Printf("  grid[%d][%d] = %d\n", i, j, val)
		}
	}

	fmt.Println()
}

// SliceSortingExample demonstrates sorting slices
func SliceSortingExample() {
	fmt.Println("=== SLICE SORTING EXAMPLE ===")

	// Sorting integers
	numbers := []int{5, 2, 9, 1, 3, 6}
	fmt.Println("Original numbers:", numbers)

	sort.Ints(numbers)
	fmt.Println("Sorted numbers:", numbers)

	// Sorting strings
	names := []string{"Charlie", "Alice", "Bob", "David"}
	fmt.Println("Original names:", names)

	sort.Strings(names)
	fmt.Println("Sorted names:", names)

	// Custom sorting with sort.Slice
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"David", 20},
	}

	// Sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("People sorted by age:", people)

	// Sort by name
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Println("People sorted by name:", people)

	fmt.Println()
}

// CommonSliceOperationsExample shows frequently used slice operations
func CommonSliceOperationsExample() {
	fmt.Println("=== COMMON SLICE OPERATIONS EXAMPLE ===")

	// Filtering elements
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Filter even numbers
	var evens []int
	for _, n := range numbers {
		if n%2 == 0 {
			evens = append(evens, n)
		}
	}
	fmt.Println("Even numbers:", evens)

	// Mapping (transforming) elements
	// Double each number
	doubled := make([]int, len(numbers))
	for i, n := range numbers {
		doubled[i] = n * 2
	}
	fmt.Println("Doubled numbers:", doubled)

	// Checking if a slice contains an element
	searchFor := 5
	contains := false
	for _, n := range numbers {
		if n == searchFor {
			contains = true
			break
		}
	}
	fmt.Printf("Slice contains %d: %t\n", searchFor, contains)

	// Finding the sum of elements
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	fmt.Println("Sum of elements:", sum)

	// Converting a string to a slice of runes
	str := "Hello, 世界"
	runes := []rune(str)
	fmt.Println("String as runes:", runes)

	// Converting back to string
	backToString := string(runes)
	fmt.Println("Back to string:", backToString)

	// Joining a slice of strings
	words := []string{"Hello", "Go", "World"}
	joined := strings.Join(words, "-")
	fmt.Println("Joined string:", joined)

	// Splitting a string into a slice
	split := strings.Split(joined, "-")
	fmt.Println("Split string:", split)

	fmt.Println()
}

// PerformanceConsiderationsExample explains performance aspects of slices
func PerformanceConsiderationsExample() {
	fmt.Println("=== PERFORMANCE CONSIDERATIONS ===")

	fmt.Println("1. Pre-allocation")
	fmt.Println("   - Use make() with capacity when size is known")
	fmt.Println("   - Avoids multiple reallocations during growth")

	fmt.Println("\n2. Avoiding Unnecessary Copies")
	fmt.Println("   - Be careful with large slices in function parameters")
	fmt.Println("   - Use pointers for large slices if modification needed")
	fmt.Println("   - Pass slice by value if no modification needed (slices are references)")

	fmt.Println("\n3. Memory Leaks")
	fmt.Println("   - Slices can cause memory leaks by holding references")
	fmt.Println("   - Use copy() to create a new slice without references")
	fmt.Println("   - Be mindful of very large backing arrays")

	fmt.Println("\n4. Slice Tricks")
	fmt.Println("   - Append is efficient due to growth strategy")
	fmt.Println("   - Filter in place for better memory efficiency")
	fmt.Println("   - Slicing doesn't copy elements, only creates a view")

	fmt.Println()
}

// ArraysAndSlicesInterviewQuestions presents common interview questions
func ArraysAndSlicesInterviewQuestions() {
	fmt.Println("=========================================")
	fmt.Println("COMMON INTERVIEW QUESTIONS:")
	fmt.Println("=========================================")

	fmt.Println("1. What is the difference between arrays and slices in Go?")
	fmt.Println("   - Arrays have fixed size, slices are dynamic")
	fmt.Println("   - Arrays are values (copied when assigned), slices are references")
	fmt.Println("   - Arrays' size is part of their type, slices' isn't")
	fmt.Println("   - Arrays are less flexible but have slightly better performance")
	fmt.Println()

	fmt.Println("2. How does slice capacity work and when does it grow?")
	fmt.Println("   - Capacity is how many elements slice can hold without reallocation")
	fmt.Println("   - When appending beyond capacity, Go creates a new backing array")
	fmt.Println("   - Growth is typically 2x the current capacity")
	fmt.Println("   - Inefficient append can lead to O(n²) operations instead of amortized O(n)")
	fmt.Println()

	fmt.Println("3. How do slices share memory, and what are the implications?")
	fmt.Println("   - Multiple slices can share the same backing array")
	fmt.Println("   - Modifying one slice can affect others that share memory")
	fmt.Println("   - Appending may cause slice to get new backing array and break sharing")
	fmt.Println("   - Use copy() to avoid unintended sharing")
	fmt.Println()

	fmt.Println("4. How would you implement a stack using slices?")
	fmt.Println("   - Push: append(stack, value)")
	fmt.Println("   - Pop: value, stack = stack[len(stack)-1], stack[:len(stack)-1]")
	fmt.Println("   - Peek: stack[len(stack)-1]")
	fmt.Println("   - IsEmpty: len(stack) == 0")
	fmt.Println()

	fmt.Println("5. What's the most efficient way to remove an element from a slice?")
	fmt.Println("   - From end: slice = slice[:len(slice)-1] - O(1), maintains order")
	fmt.Println("   - From start: slice = slice[1:] - O(1), maintains order")
	fmt.Println("   - From middle maintaining order: slice = append(slice[:i], slice[i+1:]...) - O(n)")
	fmt.Println("   - From middle not maintaining order: slice[i] = slice[len(slice)-1]; slice = slice[:len(slice)-1] - O(1)")
	fmt.Println()

	fmt.Println("6. How would you create a deep copy of a slice?")
	fmt.Println("   - Use copy(): dest := make([]T, len(src)); copy(dest, src)")
	fmt.Println("   - For slices of slices, need to loop and copy each inner slice")
	fmt.Println()

	fmt.Println("7. What happens when you pass a slice to a function?")
	fmt.Println("   - Slice header is copied (pass by value)")
	fmt.Println("   - Header contains pointer to backing array (reference semantics)")
	fmt.Println("   - Changes to elements affect the original slice")
	fmt.Println("   - Reslicing or appending might not affect the original slice")
	fmt.Println()

	fmt.Println("8. What are some common slice bugs?")
	fmt.Println("   - Out of range panics: accessing indexes beyond length")
	fmt.Println("   - Memory leaks: keeping references to small pieces of large arrays")
	fmt.Println("   - Unexpected sharing: mutations affecting unrelated code")
	fmt.Println("   - Inefficient repeated growth: not pre-allocating when size is known")
	fmt.Println()

	fmt.Println("9. How would you implement a queue with slices?")
	fmt.Println("   - Enqueue: append(queue, value)")
	fmt.Println("   - Dequeue: value, queue = queue[0], queue[1:]")
	fmt.Println("   - Note: simple implementation can be inefficient due to shifting")
	fmt.Println("   - For high-performance, use a circular buffer or linked list")
	fmt.Println()

	fmt.Println("10. How does garbage collection work with slices?")
	fmt.Println("    - The backing array is garbage collected when no slices reference it")
	fmt.Println("    - Slices that reference small parts of large arrays prevent collection")
	fmt.Println("    - Slicing very large arrays and keeping small portions can waste memory")
	fmt.Println("    - Use copy() to allow large backing arrays to be garbage collected")
	fmt.Println()
}
