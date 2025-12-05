package main

import "fmt"

// ========== FUNGSI DASAR ==========

// 1. Fungsi sederhana tanpa parameter dan return value
func sayHello() {
	fmt.Println("Hello, World!")
}

// 2. Fungsi dengan parameter
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 3. Fungsi dengan return value
func add(a, b int) int {
	return a + b
}

// 4. Fungsi dengan multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("tidak bisa dibagi dengan nol")
	}
	return a / b, nil
}

// 5. Fungsi dengan named return values
func calculate(a, b int) (sum, product int) {
	sum = a + b
	product = a * b
	return // naked return
}

// 6. Fungsi dengan variadic parameters (parameter tak terbatas)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Contoh penggunaan fungsi dasar
func DemoBasicFunctions() {
	fmt.Println("=== FUNGSI DASAR ===")
	
	// Fungsi sederhana
	sayHello()
	greet("Alice")
	
	// Fungsi dengan return
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)
	
	// Multiple return values
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", quotient)
	}
	
	// Named return values
	s, p := calculate(4, 5)
	fmt.Printf("Sum: %d, Product: %d\n", s, p)
	
	// Variadic parameters
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum of 1,2,3,4,5 = %d\n", total)
	
	fmt.Println()
}