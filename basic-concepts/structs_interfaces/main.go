package main

import (
	"fmt"
	"math"
	"strings"
)

// STRUCTS
// Struct definition
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address   Address // Embedded struct
}

// Another struct definition
type Address struct {
	Street  string
	City    string
	Country string
	ZipCode string
}

// Struct with embedded struct (composition)
type Employee struct {
	Person  // Embedded struct (anonymous field)
	Company string
	Salary  float64
}

// Struct with tags (used for serialization)
type Product struct {
	ID          int     `json:"id" xml:"product_id"`
	Name        string  `json:"name" xml:"product_name"`
	Price       float64 `json:"price" xml:"price"`
	Description string  `json:"desc,omitempty" xml:"description,omitempty"`
}

// Method with value receiver
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// Method with pointer receiver
func (p *Person) UpdateName(firstName, lastName string) {
	p.FirstName = firstName
	p.LastName = lastName
}

// INTERFACES
// Interface definition
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Implementing Shape interface with Circle struct
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Implementing Shape interface with Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Additional method specific to Rectangle
func (r Rectangle) IsSquare() bool {
	return r.Width == r.Height
}

// Empty interface
func PrintAny(any interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", any, any)
}

// Interface embedding
type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// Interface composition
type WriteCloser interface {
	Writer
	Closer
}

// Simple implementation of WriteCloser
type StringWriter struct {
	data string
}

func (sw *StringWriter) Write(p []byte) (int, error) {
	sw.data += string(p)
	return len(p), nil
}

func (sw *StringWriter) Close() error {
	sw.data = ""
	return nil
}

// Stringer interface example (built-in)
type Book struct {
	Title  string
	Author string
	Pages  int
}

// Implementing fmt.Stringer interface
func (b Book) String() string {
	return fmt.Sprintf("%s by %s (%d pages)", b.Title, b.Author, b.Pages)
}

func main() {
	fmt.Println("=== STRUCTS ===")

	// Creating a struct
	addr := Address{
		Street:  "123 Main St",
		City:    "San Francisco",
		Country: "USA",
		ZipCode: "94105",
	}

	// Creating a struct using the new keyword (returns a pointer)
	p1 := new(Person)
	p1.FirstName = "John"
	p1.LastName = "Doe"
	p1.Age = 30
	p1.Address = addr

	// Creating a struct with literal syntax
	p2 := Person{
		FirstName: "Jane",
		LastName:  "Smith",
		Age:       28,
		Address: Address{
			Street:  "456 Market St",
			City:    "New York",
			Country: "USA",
			ZipCode: "10001",
		},
	}

	// Creating struct with positional arguments (not recommended - brittle)
	p3 := Person{"Bob", "Johnson", 35, Address{"789 Oak St", "Chicago", "USA", "60601"}}

	// Omitting fields (will use zero values)
	p4 := Person{FirstName: "Alice", LastName: "Brown"}

	// Printing structs
	fmt.Println("Person 1:", p1)
	fmt.Println("Person 2:", p2)
	fmt.Println("Person 3:", p3)
	fmt.Println("Person 4:", p4)

	// Accessing struct fields
	fmt.Println("\nAccessing struct fields:")
	fmt.Println("p2 first name:", p2.FirstName)
	fmt.Println("p2 city:", p2.Address.City)

	// Calling a method on struct
	fmt.Println("\nCalling methods:")
	fmt.Println("Full name:", p2.FullName())

	// Calling a method with pointer receiver
	p2.UpdateName("Janet", "Smith-Jones")
	fmt.Println("Updated full name:", p2.FullName())

	// Struct with embedded struct
	fmt.Println("\nStruct composition:")
	emp := Employee{
		Person: Person{
			FirstName: "Mike",
			LastName:  "Wilson",
			Age:       42,
			Address: Address{
				Street:  "101 Pine St",
				City:    "Boston",
				Country: "USA",
				ZipCode: "02108",
			},
		},
		Company: "Acme Corp",
		Salary:  75000,
	}

	// Accessing fields from embedded struct
	fmt.Println("Employee name:", emp.FirstName, emp.LastName) // Direct access to Person fields
	fmt.Println("Employee full name:", emp.FullName())         // Direct access to Person methods
	fmt.Println("Employee company:", emp.Company)

	// Anonymous structs
	fmt.Println("\nAnonymous structs:")
	point := struct {
		X, Y int
	}{
		X: 10,
		Y: 20,
	}
	fmt.Println("Point:", point)

	fmt.Println("\n=== INTERFACES ===")

	// Creating shape instances
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	// Using interface type
	shapes := []Shape{circle, rectangle}

	// Iterating through interface slice
	for _, shape := range shapes {
		fmt.Printf("Shape: %#v\n", shape)
		fmt.Printf("Area: %.2f\n", shape.Area())
		fmt.Printf("Perimeter: %.2f\n", shape.Perimeter())

		// Type assertion to check specific type
		if rect, ok := shape.(Rectangle); ok {
			fmt.Printf("Is square: %t\n", rect.IsSquare())
		}

		fmt.Println()
	}

	// Type switches
	fmt.Println("Type switches:")
	for _, shape := range shapes {
		switch s := shape.(type) {
		case Circle:
			fmt.Printf("Circle with radius %.2f\n", s.Radius)
		case Rectangle:
			fmt.Printf("Rectangle with width %.2f and height %.2f\n", s.Width, s.Height)
		default:
			fmt.Println("Unknown shape")
		}
	}

	// Empty interface
	fmt.Println("\nEmpty interface examples:")
	PrintAny(42)
	PrintAny("Hello")
	PrintAny(true)
	PrintAny(circle)

	// Interface implementation check (compile-time)
	var _ Shape = Circle{}    // This will compile only if Circle implements Shape
	var _ Shape = Rectangle{} // This will compile only if Rectangle implements Shape

	// Interface composition
	fmt.Println("\nInterface composition:")
	sw := &StringWriter{}

	// Write data
	sw.Write([]byte("Hello, "))
	sw.Write([]byte("Go!"))
	fmt.Println("StringWriter data:", sw.data)

	// Use the interface
	var wc WriteCloser = sw
	wc.Write([]byte(" Welcome."))
	fmt.Println("After writing through interface:", sw.data)

	wc.Close()
	fmt.Println("After closing:", sw.data)

	// Stringer interface
	fmt.Println("\nStringer interface:")
	book := Book{
		Title:  "The Go Programming Language",
		Author: "Alan A. A. Donovan and Brian W. Kernighan",
		Pages:  380,
	}
	fmt.Println(book) // fmt.Println uses the String() method

	// Error handling example
	fmt.Println("\nError handling example:")
	if err := validateNameInput(""); err != nil {
		fmt.Println("Error:", err)
	}
	if err := validateNameInput("John"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Name validation successful")
	}

	// nil interface values
	fmt.Println("\nNil interface values:")
	var s1 Shape
	// s1.Area() would panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Println("s1 == nil:", s1 == nil)

	// Interface with nil value
	var c *Circle
	var s2 Shape = c // Non-nil interface containing nil pointer
	fmt.Println("c == nil:", c == nil)
	fmt.Println("s2 == nil:", s2 == nil) // false, because interface is not nil
}

// Implementing error interface
type InputValidationError struct {
	Field   string
	Message string
}

func (v InputValidationError) Error() string {
	return fmt.Sprintf("Validation error for %s: %s", v.Field, v.Message)
}

// Function returning a custom error
func validateNameInput(name string) error {
	if strings.TrimSpace(name) == "" {
		return InputValidationError{
			Field:   "name",
			Message: "cannot be empty",
		}
	}
	return nil
}

/*
Common interview questions about structs and interfaces in Go:

1. What is a struct in Go?
   - A composite data type that groups together variables under a single name
   - Similar to classes in other languages, but without inheritance
   - Can have methods associated with them

2. What's the difference between value receiver and pointer receiver methods?
   - Value receiver: Gets a copy of the struct, cannot modify the original
   - Pointer receiver: Gets a pointer to the struct, can modify the original
   - Value receivers are good for immutable methods, pointer receivers for mutable methods

3. What is struct embedding in Go?
   - A form of composition where one struct is embedded in another
   - Fields and methods of the embedded struct are accessible directly
   - Go's way of achieving code reuse without inheritance

4. What are struct tags in Go?
   - Metadata attached to struct fields as string literals
   - Used by reflection to provide instructions for how to handle fields
   - Common in JSON/XML serialization, ORM mapping, and validation libraries

5. What is an interface in Go?
   - A type that defines a set of methods (a method signature)
   - Only specifies what methods a type must have, not how they're implemented
   - Implemented implicitly - no "implements" keyword needed

6. What's the empty interface (interface{}) in Go?
   - An interface with no methods
   - Every type in Go implements the empty interface
   - Used when you need to handle values of unknown type

7. What's type assertion in Go?
   - A way to extract the underlying concrete value from an interface
   - Syntax: value, ok := interfaceValue.(ConcreteType)
   - If ok is false, the conversion failed

8. What's a type switch in Go?
   - A switch statement that tests the dynamic type of an interface value
   - Allows different code to run depending on the concrete type
   - Syntax: switch v := interfaceValue.(type) { case Type1: ... }

9. Can a type implement multiple interfaces in Go?
   - Yes, a type can implement any number of interfaces
   - Each interface is satisfied implicitly by implementing its methods

10. What is the difference between a nil interface and an interface with a nil value?
    - A nil interface has no type and no value (var i Interface = nil)
    - An interface with a nil value has a type but its value is nil (var p *Person = nil; var i Interface = p)
    - Calling methods on a nil interface will panic, but calling methods on an interface with a nil value is valid
*/
