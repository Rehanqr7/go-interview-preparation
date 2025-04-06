package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	fmt.Println("=========================================")
	fmt.Println("GO MAPS EXAMPLES")
	fmt.Println("=========================================")

	// Basic map examples
	BasicMapsExample()

	// Map operations
	MapOperationsExample()

	// Maps with complex keys
	ComplexKeysExample()

	// Maps of maps
	MapsOfMapsExample()

	// Concurrent map access
	ConcurrentMapAccessExample()

	// Maps with structs
	MapsWithStructsExample()

	// Maps performance considerations
	MapsPerformanceExample()

	// Common map operations
	CommonMapOperationsExample()

	// Map gotchas and tips
	MapGotchasAndTipsExample()

	// Interview questions
	MapsInterviewQuestions()
}

// BasicMapsExample demonstrates map declaration and initialization
func BasicMapsExample() {
	fmt.Println("=== BASIC MAPS EXAMPLE ===")

	// Declaring and initializing maps
	var emptyMap map[string]int // nil map (zero value)

	fmt.Println("Empty map:", emptyMap)
	fmt.Println("Empty map is nil:", emptyMap == nil)

	// Creating a map with make
	ages := make(map[string]int)

	// Add key-value pairs
	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35

	fmt.Println("Ages map:", ages)
	fmt.Println("Number of entries:", len(ages))

	// Map literal syntax
	scores := map[string]int{
		"Alice":   95,
		"Bob":     80,
		"Charlie": 90,
	}
	fmt.Println("Scores map:", scores)

	// Empty map with make
	emptyButNotNil := make(map[string]int)
	fmt.Println("Empty but not nil map:", emptyButNotNil)
	fmt.Println("Is nil?", emptyButNotNil == nil)

	// Empty map literal
	emptyLiteral := map[string]int{}
	fmt.Println("Empty literal map:", emptyLiteral)
	fmt.Println("Is nil?", emptyLiteral == nil)

	fmt.Println()
}

// MapOperationsExample demonstrates common map operations
func MapOperationsExample() {
	fmt.Println("=== MAP OPERATIONS EXAMPLE ===")

	// Create a map of fruits and their counts
	fruits := map[string]int{
		"apple":  5,
		"banana": 8,
		"orange": 3,
	}
	fmt.Println("Initial map:", fruits)

	// Retrieving values
	appleCount := fruits["apple"]
	fmt.Println("Apple count:", appleCount)

	// Retrieving values with "comma ok" idiom
	orangeCount, exists := fruits["orange"]
	fmt.Printf("Orange count: %d, exists: %t\n", orangeCount, exists)

	// Check for non-existent key
	grapeCount, exists := fruits["grape"]
	fmt.Printf("Grape count: %d, exists: %t\n", grapeCount, exists)

	// Adding new key-value pairs
	fruits["grape"] = 10
	fmt.Println("After adding grape:", fruits)

	// Updating existing entries
	fruits["apple"] = 7
	fmt.Println("After updating apple:", fruits)

	// Deleting entries
	delete(fruits, "banana")
	fmt.Println("After deleting banana:", fruits)

	// Deleting non-existent entry (no error)
	delete(fruits, "pear")
	fmt.Println("After deleting non-existent key:", fruits)

	// Zero value for non-existent keys
	fmt.Println("Requesting non-existent key:", fruits["pear"]) // Returns 0 (zero value for int)

	// Iterating over map using range (order is not guaranteed)
	fmt.Println("Iterating over map:")
	for key, value := range fruits {
		fmt.Printf("  %s: %d\n", key, value)
	}

	// Iterating over just keys
	fmt.Println("Keys in the map:")
	for key := range fruits {
		fmt.Printf("  %s\n", key)
	}

	// Clearing a map
	for k := range fruits {
		delete(fruits, k)
	}
	fmt.Println("After clearing map:", fruits)
	fmt.Println("Length of cleared map:", len(fruits))

	fmt.Println()
}

// ComplexKeysExample demonstrates using complex types as map keys
func ComplexKeysExample() {
	fmt.Println("=== COMPLEX KEYS EXAMPLE ===")

	// Integer keys
	intMap := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	fmt.Println("Map with integer keys:", intMap)

	// Boolean keys
	boolMap := map[bool]string{
		true:  "yes",
		false: "no",
	}
	fmt.Println("Map with boolean keys:", boolMap)

	// Struct as key (must be comparable)
	type Point struct {
		X, Y int
	}

	pointMap := map[Point]string{
		{1, 2}: "point A",
		{3, 4}: "point B",
	}
	fmt.Println("Map with struct keys:", pointMap)

	// Accessing values with struct keys
	p := Point{1, 2}
	fmt.Println("Value for point {1, 2}:", pointMap[p])

	// Array as key (fixed size arrays are comparable)
	arrayMap := map[[3]int]string{
		{1, 2, 3}: "array 1",
		{4, 5, 6}: "array 2",
	}
	fmt.Println("Map with array keys:", arrayMap)

	// NOT ALLOWED: Slices as keys (uncommenting will cause error)
	// sliceMap := map[[]int]string{} // Compile error: invalid map key type []int

	// NOT ALLOWED: Maps as keys (uncommenting will cause error)
	// mapMap := map[map[string]int]string{} // Compile error: invalid map key type map[string]int

	// Using string keys with struct values
	type Person struct {
		Name string
		Age  int
	}

	people := map[string]Person{
		"alice":   {"Alice Smith", 30},
		"bob":     {"Bob Johnson", 25},
		"charlie": {"Charlie Brown", 35},
	}
	fmt.Println("Map with struct values:", people)

	// Accessing struct fields
	fmt.Printf("Charlie's age: %d\n", people["charlie"].Age)

	fmt.Println()
}

// MapsOfMapsExample demonstrates nested maps
func MapsOfMapsExample() {
	fmt.Println("=== MAPS OF MAPS EXAMPLE ===")

	// Declaring a map of maps
	cityPopulationByCountry := map[string]map[string]int{
		"USA": {
			"New York":    8804190,
			"Los Angeles": 3898747,
			"Chicago":     2746388,
		},
		"Japan": {
			"Tokyo":    13960000,
			"Yokohama": 3760000,
			"Osaka":    2691000,
		},
	}

	fmt.Println("City populations by country:", cityPopulationByCountry)

	// Accessing nested maps
	fmt.Println("Population of New York:", cityPopulationByCountry["USA"]["New York"])

	// You must initialize inner maps before using them
	cityPopulationByCountry["Canada"] = make(map[string]int)
	cityPopulationByCountry["Canada"]["Toronto"] = 2930000
	cityPopulationByCountry["Canada"]["Vancouver"] = 675218

	fmt.Println("After adding Canada:", cityPopulationByCountry)

	// Accessing a non-existent nested key safely
	_, exists := cityPopulationByCountry["UK"]
	if !exists {
		fmt.Println("No data for UK")
	}

	// This would panic if UK didn't exist and we didn't check:
	// population := cityPopulationByCountry["UK"]["London"] // Panic: assignment to entry in nil map

	// Safer way to access nested map values
	population, err := getNestedMapValue(cityPopulationByCountry, "Japan", "Tokyo")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Population of Tokyo:", population)
	}

	// Try a non-existent path
	population, err = getNestedMapValue(cityPopulationByCountry, "UK", "London")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Population of London:", population)
	}

	fmt.Println()
}

// Helper function to safely get values from nested maps
func getNestedMapValue(m map[string]map[string]int, outerKey, innerKey string) (int, error) {
	if innerMap, ok := m[outerKey]; ok {
		if val, ok := innerMap[innerKey]; ok {
			return val, nil
		}
		return 0, fmt.Errorf("inner key %q not found", innerKey)
	}
	return 0, fmt.Errorf("outer key %q not found", outerKey)
}

// ConcurrentMapAccessExample demonstrates thread-safe map access
func ConcurrentMapAccessExample() {
	fmt.Println("=== CONCURRENT MAP ACCESS EXAMPLE ===")

	// Regular maps are not safe for concurrent use
	// This would cause a race condition if run with -race flag:
	/*
		counters := make(map[string]int)
		var wg sync.WaitGroup

		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				counters["key"] += 1 // Race condition!
			}()
		}

		wg.Wait()
	*/

	// Using sync.RWMutex for safe concurrent access
	type SafeMap struct {
		mu   sync.RWMutex
		data map[string]int
	}

	safeMap := SafeMap{
		data: make(map[string]int),
	}

	// Safe write function
	increment := func(key string) {
		safeMap.mu.Lock()
		defer safeMap.mu.Unlock()
		safeMap.data[key]++
	}

	// Safe read function
	get := func(key string) int {
		safeMap.mu.RLock()
		defer safeMap.mu.RUnlock()
		return safeMap.data[key]
	}

	// Concurrent increments
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment("counter")
		}()
	}

	wg.Wait()
	fmt.Println("Final count:", get("counter"))

	// Using sync.Map for concurrent access
	var syncMap sync.Map

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n%10) // Use 10 different keys

			// Load or initialize
			value, _ := syncMap.LoadOrStore(key, 0)
			// Increment and store
			syncMap.Store(key, value.(int)+1)
		}(i)
	}

	wg.Wait()

	// Print contents of sync.Map
	fmt.Println("sync.Map contents:")
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("  %v: %v\n", key, value)
		return true // Continue iteration
	})

	fmt.Println()
}

// MapsWithStructsExample demonstrates maps with struct values
func MapsWithStructsExample() {
	fmt.Println("=== MAPS WITH STRUCTS EXAMPLE ===")

	// Define a struct type
	type Employee struct {
		Name     string
		ID       int
		Position string
		Salary   float64
	}

	// Create a map with Employee values
	employees := map[string]Employee{
		"alice": {
			Name:     "Alice Smith",
			ID:       1001,
			Position: "Software Engineer",
			Salary:   90000,
		},
		"bob": {
			Name:     "Bob Johnson",
			ID:       1002,
			Position: "Product Manager",
			Salary:   95000,
		},
	}

	fmt.Println("Employee map:", employees)

	// Accessing struct fields
	fmt.Printf("Alice's position: %s\n", employees["alice"].Position)

	// Cannot modify struct fields directly through map
	// This would cause a compile error:
	// employees["alice"].Salary += 5000 // Cannot assign to employees["alice"].Salary

	// Instead, you need to replace the entire struct:
	alice := employees["alice"]
	alice.Salary += 5000
	employees["alice"] = alice

	fmt.Printf("Alice's new salary: %.2f\n", employees["alice"].Salary)

	// Using maps with pointers to structs allows direct modification
	type Department struct {
		Name     string
		Location string
		Budget   float64
	}

	// Map with pointers to structs
	departments := map[string]*Department{
		"engineering": {
			Name:     "Engineering",
			Location: "Building A",
			Budget:   1000000,
		},
		"marketing": {
			Name:     "Marketing",
			Location: "Building B",
			Budget:   500000,
		},
	}

	// Now we can modify fields directly
	departments["engineering"].Budget += 200000

	fmt.Println("Engineering department budget:",
		departments["engineering"].Budget)

	// Adding a new department
	departments["sales"] = &Department{
		Name:     "Sales",
		Location: "Building C",
		Budget:   750000,
	}

	// Printing the departments
	fmt.Println("Departments:")
	for name, dept := range departments {
		fmt.Printf("  %s: %s, Budget: %.2f\n",
			name, dept.Location, dept.Budget)
	}

	fmt.Println()
}

// MapsPerformanceExample discusses map performance considerations
func MapsPerformanceExample() {
	fmt.Println("=== MAPS PERFORMANCE EXAMPLE ===")

	fmt.Println("Map Performance Considerations:")

	fmt.Println("\n1. Initial Capacity")
	fmt.Println("   - Use make(map[K]V, cap) to pre-allocate for known size")
	fmt.Println("   - Reduces rehashing operations for better performance")
	fmt.Println("   - Example: users := make(map[string]User, 10000)")

	fmt.Println("\n2. Map Growth")
	fmt.Println("   - Maps grow automatically as needed")
	fmt.Println("   - Growth triggers rehashing (moving all items)")
	fmt.Println("   - Pre-allocation helps avoid this cost")

	fmt.Println("\n3. Key Types and Performance")
	fmt.Println("   - Simple keys (int, string) typically perform better")
	fmt.Println("   - Complex keys (structs) require more computation for hashing")
	fmt.Println("   - Comparable types only (no slices, maps, functions)")

	fmt.Println("\n4. Access Complexity")
	fmt.Println("   - Average case O(1) for lookups, insertions, deletions")
	fmt.Println("   - Worst case O(n) if many hash collisions")
	fmt.Println("   - Go's implementation uses high-quality hash functions")

	fmt.Println("\n5. Memory Usage")
	fmt.Println("   - Maps have higher memory overhead than arrays or slices")
	fmt.Println("   - Overhead is due to the hash table structure")
	fmt.Println("   - Consider this for memory-sensitive applications")

	fmt.Println()
}

// CommonMapOperationsExample demonstrates frequently used map patterns
func CommonMapOperationsExample() {
	fmt.Println("=== COMMON MAP OPERATIONS EXAMPLE ===")

	// Counting occurrences
	words := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}
	fmt.Println("Word counts:", counts)

	// Set implementation (map with empty struct values)
	uniqueWords := map[string]struct{}{}
	for _, word := range words {
		uniqueWords[word] = struct{}{}
	}

	fmt.Println("Unique words:")
	for word := range uniqueWords {
		fmt.Printf("  %s\n", word)
	}

	// Getting sorted keys from a map
	fruits := map[string]int{
		"apple":  5,
		"banana": 8,
		"orange": 3,
		"grape":  10,
	}

	keys := make([]string, 0, len(fruits))
	for k := range fruits {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println("Sorted fruits by name:")
	for _, k := range keys {
		fmt.Printf("  %s: %d\n", k, fruits[k])
	}

	// Grouping with maps
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 30},
		{"David", 25},
		{"Eve", 35},
	}

	byAge := make(map[int][]string)
	for _, person := range people {
		byAge[person.Age] = append(byAge[person.Age], person.Name)
	}

	fmt.Println("People grouped by age:")
	for age, names := range byAge {
		fmt.Printf("  Age %d: %v\n", age, names)
	}

	// Inverting a map (value -> key)
	scores := map[string]int{
		"Alice":   95,
		"Bob":     80,
		"Charlie": 90,
	}

	nameByScore := make(map[int]string)
	for name, score := range scores {
		nameByScore[score] = name
	}

	fmt.Println("Names by score:", nameByScore)
	fmt.Println("Note: This only works if values are unique!")

	// Merging maps
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 3, "c": 4}

	merged := make(map[string]int)
	for k, v := range map1 {
		merged[k] = v
	}
	for k, v := range map2 {
		merged[k] = v // Overwrites if key exists
	}

	fmt.Println("Merged map:", merged)

	// Finding all keys with a particular value
	colorMap := map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"grape":  "purple",
		"cherry": "red",
		"lemon":  "yellow",
	}

	redFruits := []string{}
	for fruit, color := range colorMap {
		if color == "red" {
			redFruits = append(redFruits, fruit)
		}
	}

	fmt.Println("Red fruits:", redFruits)

	fmt.Println()
}

// MapGotchasAndTipsExample demonstrates common pitfalls and best practices
func MapGotchasAndTipsExample() {
	fmt.Println("=== MAP GOTCHAS AND TIPS EXAMPLE ===")

	fmt.Println("1. Nil map access")
	fmt.Println("   - Reading from nil map returns zero values")
	fmt.Println("   - Writing to nil map causes panic")

	var nilMap map[string]int
	fmt.Println("   Read from nil map:", nilMap["key"]) // Returns 0
	// Uncommenting this line would panic:
	// nilMap["key"] = 1 // panic: assignment to entry in nil map

	fmt.Println("\n2. Map iteration order")
	fmt.Println("   - Map iteration order is not guaranteed")
	fmt.Println("   - Order may change between runs or even iterations")
	fmt.Println("   - Sort keys first if order matters")

	fmt.Println("\n3. Map equality")
	fmt.Println("   - Maps can't be compared with == except with nil")
	fmt.Println("   - Must write custom equality function")

	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"a": 1, "b": 2}
	fmt.Println("   Maps equal?", mapEqual(map1, map2))

	fmt.Println("\n4. Zero value behavior")
	fmt.Println("   - Accessing non-existent key returns zero value")
	fmt.Println("   - Use 'comma ok' idiom to check existence")

	users := map[string]int{"alice": 30}
	age := users["bob"] // Returns 0 (zero value for int)
	fmt.Println("   Bob's age (non-existent key):", age)

	age, exists := users["bob"]
	fmt.Printf("   Bob's age: %d, exists: %t\n", age, exists)

	fmt.Println("\n5. Map capacity")
	fmt.Println("   - No direct way to get current capacity")
	fmt.Println("   - No way to shrink a map once it grows")
	fmt.Println("   - Only way is to create a new map")

	fmt.Println("\n6. Concurrent access")
	fmt.Println("   - Maps are not safe for concurrent access")
	fmt.Println("   - Use sync.RWMutex or sync.Map for concurrent use")

	fmt.Println()
}

// Helper function to compare two maps
func mapEqual(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}

	return true
}

// MapsInterviewQuestions presents common interview questions about maps
func MapsInterviewQuestions() {
	fmt.Println("=========================================")
	fmt.Println("COMMON INTERVIEW QUESTIONS:")
	fmt.Println("=========================================")

	fmt.Println("1. What is a map in Go and how does it work internally?")
	fmt.Println("   - Map is a hash table implementation")
	fmt.Println("   - Stores key-value pairs with O(1) average lookup")
	fmt.Println("   - Implemented as hash table with buckets for collisions")
	fmt.Println("   - Uses a high-quality hash function to minimize collisions")
	fmt.Println()

	fmt.Println("2. What types can be used as map keys in Go?")
	fmt.Println("   - Only comparable types: bool, numeric, string, pointer, channel")
	fmt.Println("   - Arrays and structs if their elements are comparable")
	fmt.Println("   - Cannot use: slices, maps, functions (non-comparable)")
	fmt.Println()

	fmt.Println("3. How do you check if a key exists in a map?")
	fmt.Println("   - Use the comma ok idiom: value, ok := map[key]")
	fmt.Println("   - ok is true if key exists, false otherwise")
	fmt.Println("   - value will be the zero value if key doesn't exist")
	fmt.Println()

	fmt.Println("4. What's the time complexity of map operations in Go?")
	fmt.Println("   - Average case: O(1) for lookup, insert, delete")
	fmt.Println("   - Worst case: O(n) if many hash collisions")
	fmt.Println("   - Iteration: O(n) where n is the number of entries")
	fmt.Println()

	fmt.Println("5. How to make maps safe for concurrent use?")
	fmt.Println("   - Option 1: Use sync.RWMutex around map operations")
	fmt.Println("   - Option 2: Use sync.Map for specific concurrent patterns")
	fmt.Println("   - Option 3: Use channels to coordinate access")
	fmt.Println()

	fmt.Println("6. What happens when you access a key that doesn't exist?")
	fmt.Println("   - Returns zero value for the value type")
	fmt.Println("   - Does not panic or return an error")
	fmt.Println("   - Use comma ok idiom to check existence")
	fmt.Println()

	fmt.Println("7. How to implement a set in Go?")
	fmt.Println("   - Use map with empty struct values: map[KeyType]struct{}")
	fmt.Println("   - Empty struct takes 0 bytes of memory")
	fmt.Println("   - Check membership with: _, exists := set[item]")
	fmt.Println()

	fmt.Println("8. What's the difference between delete(map, key) and map[key] = Zero?")
	fmt.Println("   - delete removes entry completely, may free memory")
	fmt.Println("   - setting to zero value keeps entry with zero value")
	fmt.Println("   - comma ok idiom will return different results")
	fmt.Println()

	fmt.Println("9. How to get a map's entries in a specific order?")
	fmt.Println("   - Extract keys to a slice")
	fmt.Println("   - Sort the keys slice")
	fmt.Println("   - Iterate over sorted keys and access map")
	fmt.Println()

	fmt.Println("10. Can a map be used as a key in another map?")
	fmt.Println("    - No, maps are not comparable")
	fmt.Println("    - Would need to convert map to a comparable representation")
	fmt.Println("    - Could use encoding/json or a custom string representation")
	fmt.Println()
}
