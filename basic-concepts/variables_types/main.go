package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== VARIABLES AND TYPES ===")

	// Basic variable examples
	var name string = "John"
	var age int = 30
	var salary float64 = 50000.50
	var employed bool = true

	// Type inference with :=
	department := "Engineering"
	level := 2

	fmt.Println("Basic variable examples:")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)
	fmt.Println("Employed:", employed)
	fmt.Println("Department:", department)
	fmt.Println("Level:", level)

	// Basic types demonstration
	fmt.Println("\nBasic types in Go:")

	var integer int = 42
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647 // int32 is the same as rune
	var integer64 int64 = 9223372036854775807
	var unsignedInt uint = 42
	var floatingPoint float64 = 3.14159
	var complexNum complex128 = 3 + 4i
	var boolean bool = true
	var str string = "Hello, Go!"

	fmt.Printf("int: %d, type: %T\n", integer, integer)
	fmt.Printf("int8: %d, type: %T\n", integer8, integer8)
	fmt.Printf("int16: %d, type: %T\n", integer16, integer16)
	fmt.Printf("int32 (rune): %d, type: %T\n", integer32, integer32)
	fmt.Printf("int64: %d, type: %T\n", integer64, integer64)
	fmt.Printf("uint: %d, type: %T\n", unsignedInt, unsignedInt)
	fmt.Printf("float64: %g, type: %T\n", floatingPoint, floatingPoint)
	fmt.Printf("complex128: %v, type: %T\n", complexNum, complexNum)
	fmt.Printf("bool: %t, type: %T\n", boolean, boolean)
	fmt.Printf("string: %s, type: %T\n", str, str)

	// Constants
	const Pi = 3.14159
	const Username = "admin"
	const MaxUsers = 1000

	fmt.Println("\nConstants:")
	fmt.Println("Pi:", Pi)
	fmt.Println("Username:", Username)
	fmt.Println("MaxUsers:", MaxUsers)

	// iota example
	fmt.Println("\niota example:")

	// Days of the week using iota
	const (
		Monday    = iota // 0
		Tuesday          // 1
		Wednesday        // 2
		Thursday         // 3
		Friday           // 4
		Saturday         // 5
		Sunday           // 6
	)

	fmt.Println("Monday:", Monday)
	fmt.Println("Sunday:", Sunday)

	// Math operations
	fmt.Println("\nMath operations:")
	fmt.Println("MaxInt32:", math.MaxInt32)
	fmt.Println("Pi:", math.Pi)
	fmt.Println("Square root of 16:", math.Sqrt(16))
}

/*
Common interview questions about Go variables and types:

1. What are the zero values for different types in Go?
   - int, float: 0
   - bool: false
   - string: "" (empty string)
   - pointer, slice, map, channel, function, interface: nil

2. What's the difference between var x = 5 and x := 5?
   - var x = 5 can be used anywhere
   - x := 5 is shorthand that can only be used inside functions

3. How do you declare constants in Go?
   - Using the const keyword: const Pi = 3.14159

4. What is iota in Go?
   - iota is a counter used within constant declarations that auto-increments

5. Are Go variables case-sensitive?
   - Yes, age and Age are different variables

6. What types are available in Go?
   - Basic types: bool, string, int, float, complex
   - Composite types: array, slice, map, struct
   - Reference types: pointer, function, channel
   - Interface types
*/
