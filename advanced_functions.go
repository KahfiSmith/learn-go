package main

import "fmt"

// ========== FUNGSI DENGAN VARIADIC PARAMETERS ==========

// Fungsi sebagai variable
func mathOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// ========== CLOSURE ==========

// Closure - fungsi yang menangkap variable dari scope luar
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Closure dengan parameter
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// Higher-order function yang mengembalikan fungsi
func createValidator(min, max int) func(int) bool {
	return func(value int) bool {
		return value >= min && value <= max
	}
}

// Function yang menerima function sebagai parameter
func applyToSlice(numbers []int, fn func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = fn(num)
	}
	return result
}

// Contoh penggunaan advanced functions
func DemoAdvancedFunctions() {
	fmt.Println("=== ADVANCED FUNCTIONS ===")
	
	// Function as variable
	fmt.Println("1. Function as Variable:")
	multiply := func(a, b int) int { return a * b }
	subtract := func(a, b int) int { return a - b }
	
	result1 := mathOperation(4, 5, multiply)
	result2 := mathOperation(10, 3, subtract)
	fmt.Printf("4 * 5 = %d\n", result1)
	fmt.Printf("10 - 3 = %d\n", result2)
	
	// Closure
	fmt.Println("\n2. Closure:")
	incrementer := counter()
	fmt.Printf("Counter: %d\n", incrementer())
	fmt.Printf("Counter: %d\n", incrementer())
	fmt.Printf("Counter: %d\n", incrementer())
	
	// Closure dengan parameter
	fmt.Println("\n3. Closure dengan Parameter:")
	double := multiplier(2)
	triple := multiplier(3)
	fmt.Printf("Double 5: %d\n", double(5))
	fmt.Printf("Triple 4: %d\n", triple(4))
	
	// Function validator
	fmt.Println("\n4. Function Validator:")
	isValidAge := createValidator(0, 120)
	isValidScore := createValidator(0, 100)
	fmt.Printf("Age 25 valid? %t\n", isValidAge(25))
	fmt.Printf("Age 150 valid? %t\n", isValidAge(150))
	fmt.Printf("Score 85 valid? %t\n", isValidScore(85))
	
	// Apply function to slice
	fmt.Println("\n5. Apply Function to Slice:")
	numbers := []int{1, 2, 3, 4, 5}
	squared := applyToSlice(numbers, func(x int) int { return x * x })
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Squared: %v\n", squared)
	
	fmt.Println()
}