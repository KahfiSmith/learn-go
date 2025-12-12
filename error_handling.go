package main

import (
	"fmt"
	"log"
	"os"
)

// ========== DEFER, PANIC, RECOVER ==========

// Fungsi dengan defer
func fileOperation() {
	fmt.Println("Opening file")
	defer fmt.Println("Closing file") // akan dieksekusi terakhir
	defer fmt.Println("Cleaning up resources") // LIFO order
	fmt.Println("Processing file")
}

// Fungsi dengan multiple defer
func multipleDefer() {
	fmt.Println("Function start")
	defer fmt.Println("Defer 1") // Terakhir dieksekusi
	defer fmt.Println("Defer 2") // Kedua dieksekusi
	defer fmt.Println("Defer 3") // Pertama dieksekusi
	fmt.Println("Function end")
}

// Fungsi dengan defer dalam loop
func deferInLoop() {
	fmt.Println("Defer in loop example:")
	for i := 1; i <= 3; i++ {
		defer fmt.Printf("Deferred: %d\n", i)
	}
	fmt.Println("Loop completed")
}

// ========== PANIC DAN RECOVER ==========

// Fungsi dengan panic dan recover
func safeDivision(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
			result = 0
		}
	}()
	
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

// Fungsi yang mendemonstrasikan recover
func recoverDemo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in recoverDemo: %v\n", r)
		}
	}()
	
	fmt.Println("About to panic")
	panic("Something went wrong!")
	fmt.Println("This line will not be executed")
}

// Fungsi dengan nested recovery
func nestedRecovery() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Outer recovery: %v\n", r)
		}
	}()
	
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Inner recovery: %v\n", r)
				panic("Re-panic from inner function")
			}
		}()
		panic("Inner panic")
	}()
}

// Fungsi untuk validasi input dengan panic
func validateInput(age int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Validation failed: %v\n", r)
		}
	}()
	
	if age < 0 {
		panic("age cannot be negative")
	}
	if age > 150 {
		panic("age cannot be more than 150")
	}
	fmt.Printf("Valid age: %d\n", age)
}

// ========== REAL-WORLD EXAMPLES ==========

// Fungsi untuk membuka file dengan proper error handling
func readFileWithDefer(filename string) {
	fmt.Printf("Attempting to open file: %s\n", filename)
	
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	
	// Defer close untuk memastikan file selalu ditutup
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		} else {
			fmt.Printf("File %s closed successfully\n", filename)
		}
	}()
	
	// Simulasi pemrosesan file
	fmt.Printf("Processing file: %s\n", filename)
}

// Fungsi dengan logging dan recovery
func criticalOperation(id int) (result string) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Critical operation %d failed: %v", id, r)
			result = "FAILED"
		}
	}()
	
	fmt.Printf("Starting critical operation %d\n", id)
	
	// Simulasi operasi yang bisa panic
	if id%2 == 0 {
		panic(fmt.Sprintf("operation %d failed unexpectedly", id))
	}
	
	result = fmt.Sprintf("Operation %d completed successfully", id)
	return
}

// Fungsi dengan cleanup resources
func resourceManagement() {
	fmt.Println("Allocating resources...")
	
	// Simulasi alokasi resource
	resource1 := "Database Connection"
	resource2 := "File Handle"
	resource3 := "Network Socket"
	
	defer func() {
		fmt.Printf("Releasing: %s\n", resource3)
	}()
	
	defer func() {
		fmt.Printf("Releasing: %s\n", resource2)
	}()
	
	defer func() {
		fmt.Printf("Releasing: %s\n", resource1)
	}()
	
	fmt.Println("Using resources...")
	
	// Simulasi error yang mungkin terjadi
	if true { // Ganti dengan kondisi error
		fmt.Println("An error occurred during processing")
		return // defer functions akan tetap dieksekusi
	}
	
	fmt.Println("Processing completed successfully")
}

// Contoh penggunaan error handling functions
func DemoErrorHandling() {
	fmt.Println("=== ERROR HANDLING FUNCTIONS ===")
	
	// Defer examples
	fmt.Println("1. Defer Examples:")
	fileOperation()
	
	fmt.Println("\n2. Multiple Defer (LIFO order):")
	multipleDefer()
	
	fmt.Println("\n3. Defer in Loop:")
	deferInLoop()
	
	// Panic and recover
	fmt.Println("\n4. Panic and Recover:")
	result1 := safeDivision(10, 2)
	fmt.Printf("10 / 2 = %d\n", result1)
	
	result2 := safeDivision(10, 0)
	fmt.Printf("10 / 0 = %d\n", result2)
	
	// Recover demo
	fmt.Println("\n5. Recover Demo:")
	recoverDemo()
	fmt.Println("Program continues after recovery")
	
	// Nested recovery
	fmt.Println("\n6. Nested Recovery:")
	nestedRecovery()
	
	// Input validation
	fmt.Println("\n7. Input Validation with Panic/Recover:")
	validateInput(25)  // Valid
	validateInput(-5)  // Invalid
	validateInput(200) // Invalid
	
	// File operations with defer
	fmt.Println("\n8. File Operations with Defer:")
	readFileWithDefer("example.txt")
	
	// Critical operations
	fmt.Println("\n9. Critical Operations with Recovery:")
	for i := 1; i <= 4; i++ {
		result := criticalOperation(i)
		fmt.Printf("Result: %s\n", result)
	}
	
	// Resource management
	fmt.Println("\n10. Resource Management:")
	resourceManagement()
	
	fmt.Println()
}