package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== FUNCTIONS ===")

	// Basic function call
	fmt.Println("\n--- Basic Function ---")
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)

	// Multiple return values
	fmt.Println("\n--- Multiple Return Values ---")
	sum, difference := addAndSubtract(10, 5)
	fmt.Println("Sum:", sum, "Difference:", difference)

	// Named return values
	fmt.Println("\n--- Named Return Values ---")
	area, perimeter := rectangleProperties(5, 3)
	fmt.Println("Area:", area, "Perimeter:", perimeter)

	// Variadic function
	fmt.Println("\n--- Variadic Function ---")
	fmt.Println("Sum of numbers:", sumNumbers(1, 2, 3, 4, 5))

	// Passing a slice to a variadic function
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Sum of slice:", sumNumbers(numbers...))

	// Functions as values
	fmt.Println("\n--- Functions as Values ---")
	operation := add // Assign function to a variable
	fmt.Println("Operation result:", operation(10, 5))

	// Function as parameter
	fmt.Println("\n--- Function as Parameter ---")
	fmt.Println("Apply operation (add):", applyOperation(10, 5, add))
	fmt.Println("Apply operation (multiply):", applyOperation(10, 5, multiply))

	// Anonymous function
	fmt.Println("\n--- Anonymous Function ---")
	func(x, y int) {
		fmt.Println("Anonymous function result:", x*y)
	}(5, 3)

	// Closure (function that captures variables)
	fmt.Println("\n--- Closure ---")
	counter := createCounter()
	fmt.Println("Counter:", counter()) // 1
	fmt.Println("Counter:", counter()) // 2
	fmt.Println("Counter:", counter()) // 3

	// Another closure example
	fmt.Println("\n--- Closure with Parameter ---")
	addFive := createAdder(5)
	addTen := createAdder(10)
	fmt.Println("Add 5 to 10:", addFive(10)) // 15
	fmt.Println("Add 10 to 20:", addTen(20)) // 30

	// Higher-order function (returns a function)
	fmt.Println("\n--- Higher-Order Function ---")
	squareFunc := powerFunction(2)
	cubeFunc := powerFunction(3)
	fmt.Println("Square of 4:", squareFunc(4)) // 16
	fmt.Println("Cube of 3:", cubeFunc(3))     // 27

	// Function with deferred call
	fmt.Println("\n--- Function with Deferred Call ---")
	functionWithDefer()

	// Function with error return
	fmt.Println("\n--- Function with Error Return ---")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 0 =", result)
	}

	// Method (function attached to a type)
	fmt.Println("\n--- Method ---")
	p := person{firstName: "John", lastName: "Doe", age: 30}
	fmt.Println("Full name:", p.fullName())
	p.increaseAge(5)
	fmt.Println("New age:", p.age)

	// Function with callbacks
	fmt.Println("\n--- Function with Callbacks ---")
	inputStrings := []string{"hello", "world", "go", "programming"}

	// Example 1: Convert to uppercase
	uppercaseStrings := processStrings(inputStrings, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Println("Uppercase strings:", uppercaseStrings)

	// Example 2: Add prefix
	prefixedStrings := processStrings(inputStrings, func(s string) string {
		return "prefix_" + s
	})
	fmt.Println("Prefixed strings:", prefixedStrings)

	// Example 3: Reverse strings
	reversedStrings := processStrings(inputStrings, func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	})
	fmt.Println("Reversed strings:", reversedStrings)
}

// Basic function
func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

// Multiple return values
func addAndSubtract(a, b int) (int, int) {
	return a + b, a - b
}

// Named return values
func rectangleProperties(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // "naked" return uses the named return values
}

// Variadic function (variable number of arguments)
func sumNumbers(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function that takes a function as a parameter
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Closure (function that captures variables from its environment)
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Another closure example
func createAdder(base int) func(int) int {
	return func(x int) int {
		return base + x
	}
}

// Higher-order function (returns a function)
func powerFunction(exponent int) func(int) int {
	return func(base int) int {
		result := 1
		for i := 0; i < exponent; i++ {
			result *= base
		}
		return result
	}
}

// Function with deferred call
func functionWithDefer() {
	defer fmt.Println("This is executed last")
	fmt.Println("This is executed first")
}

// Function with error return
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Method (function attached to a type)
type person struct {
	firstName string
	lastName  string
	age       int
}

// Value receiver method
func (p person) fullName() string {
	return p.firstName + " " + p.lastName
}

// Pointer receiver method
func (p *person) increaseAge(years int) {
	p.age += years
}

// Function with callbacks
func processStrings(strings []string, callback func(string) string) []string {
	result := make([]string, len(strings))
	for i, str := range strings {
		result[i] = callback(str)
	}
	return result
}

/*
Common interview questions about Go functions:

1. What is the difference between a value receiver and a pointer receiver in methods?
   - Value receiver: Gets a copy of the value, cannot modify the original
   - Pointer receiver: Gets a pointer to the value, can modify the original

2. What are variadic functions in Go?
   - Functions that accept a variable number of arguments
   - Defined with ... before the type
   - Within the function, args become a slice of the specified type

3. Can Go return multiple values from a function?
   - Yes, Go can return multiple values, which is useful for returning both result and error

4. What are named return values in Go?
   - Return values can be named in the function signature
   - They are initialized to zero values at the start of the function
   - Can be used with a "naked return" (just "return" without arguments)

5. How do closures work in Go?
   - Functions can be created inside functions and returned
   - They "close over" variables from the outer function's scope
   - The inner function maintains access to these variables even after the outer function returns

6. What is the purpose of defer in Go functions?
   - Defers execution of a function until the surrounding function returns
   - Useful for cleanup operations (closing files, etc.)
   - Executed in LIFO order (last-in, first-out)

7. How do you handle errors in Go functions?
   - By convention, errors are returned as the last return value
   - Caller checks if error is nil to determine success
   - Error handling is explicit, not through exceptions

8. What is panic and recover in Go?
   - panic: Stops normal execution of the current goroutine
   - recover: Used inside deferred functions to regain control after a panic
   - Similar to try/catch in other languages, but meant for exceptional cases only
*/
