package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== CONTROL FLOW ===")

	// IF statements
	fmt.Println("\n--- IF Statements ---")

	age := 20
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// If with short statement
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D or below")
	}

	// FOR loops
	fmt.Println("\n--- FOR Loops ---")

	// Standard for loop
	fmt.Println("Standard for loop:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// While-like for loop
	fmt.Println("\nWhile-like for loop:")
	counter := 1
	for counter <= 5 {
		fmt.Printf("%d ", counter)
		counter++
	}
	fmt.Println()

	// Infinite loop with break
	fmt.Println("\nInfinite loop with break:")
	count := 1
	for {
		fmt.Printf("%d ", count)
		count++
		if count > 5 {
			break
		}
	}
	fmt.Println()

	// For loop with continue
	fmt.Println("\nFor loop with continue (skipping even numbers):")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// For-range over array
	fmt.Println("\nFor-range over array:")
	numbers := [5]int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("numbers[%d] = %d\n", index, value)
	}

	// For-range over string (iterates over runes)
	fmt.Println("\nFor-range over string:")
	for index, char := range "Hello" {
		fmt.Printf("[%d]: %c\n", index, char)
	}

	// SWITCH statements
	fmt.Println("\n--- SWITCH Statements ---")

	// Basic switch
	fmt.Println("Basic switch:")
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekend")
	}

	// Switch with multiple cases
	fmt.Println("\nSwitch with multiple cases:")
	fruit := "apple"
	switch fruit {
	case "apple", "pear", "banana":
		fmt.Println("Common fruit")
	case "dragonfruit", "starfruit":
		fmt.Println("Exotic fruit")
	default:
		fmt.Println("Unknown fruit")
	}

	// Switch with no expression (like if-else chain)
	fmt.Println("\nSwitch with no expression:")
	temperature := 75
	switch {
	case temperature < 32:
		fmt.Println("Freezing")
	case temperature < 50:
		fmt.Println("Cold")
	case temperature < 70:
		fmt.Println("Cool")
	case temperature < 90:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}

	// Switch with fallthrough
	fmt.Println("\nSwitch with fallthrough:")
	num := 5
	switch num {
	case 5:
		fmt.Println("Five")
		fallthrough
	case 4:
		fmt.Println("Four")
		fallthrough
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Unknown")
	}

	// DEFER Statement
	fmt.Println("\n--- DEFER Statement ---")

	// Basic defer
	fmt.Println("Basic defer example:")
	defer fmt.Println("This prints last")
	fmt.Println("This prints first")
	fmt.Println("This prints second")

	// Multiple defers (LIFO order)
	fmt.Println("\nMultiple defers (LIFO order):")
	for i := 1; i <= 3; i++ {
		defer fmt.Printf("Deferred %d\n", i)
	}

	// Defer with function call
	fmt.Println("\nDefer with function call:")
	defer printTime("End time")
	printTime("Start time")

	// Defer with arguments evaluated at defer time
	a := 10
	defer fmt.Printf("\nDeferred value of a: %d\n", a)
	a = 20
	fmt.Printf("Current value of a: %d\n", a)
}

func printTime(label string) {
	fmt.Printf("%s: %s\n", label, time.Now().Format(time.RFC3339))
}

/*
Common interview questions about Go control flow:

1. How is Go's switch statement different from C/C++/Java?
   - In Go, switch cases break automatically (no fall-through)
   - You need to use the fallthrough keyword explicitly
   - Switch cases can be expressions
   - Switch can be used without an expression (like if-else)

2. What's the difference between break and continue?
   - break exits the loop entirely
   - continue skips the rest of the current iteration and moves to the next

3. How many ways can you write a for loop in Go?
   - Standard for loop: for i := 0; i < 10; i++ {}
   - While-like loop: for condition {}
   - Infinite loop: for {}
   - Range-based: for index, value := range collection {}

4. What is defer in Go and when is it useful?
   - Defer postpones execution of a function until surrounding function returns
   - Useful for cleanup operations like closing files
   - Deferred functions are executed in LIFO order (last-in, first-out)
   - The arguments to a deferred function are evaluated when the defer statement is executed, not when the function is called

5. What is the scope of a variable declared in an if statement?
   - It's limited to the if block and any else blocks
*/
