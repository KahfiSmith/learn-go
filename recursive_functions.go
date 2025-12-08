package main

import "fmt"

// ========== RECURSIVE FUNCTION ==========

// Fungsi rekursif - factorial
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Fungsi rekursif - fibonacci
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Fungsi rekursif - menghitung pangkat
func power(base, exp int) int {
	if exp == 0 {
		return 1
	}
	if exp == 1 {
		return base
	}
	return base * power(base, exp-1)
}

// Fungsi rekursif - greatest common divisor (GCD)
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Fungsi rekursif - sum array
func sumArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	return arr[0] + sumArray(arr[1:])
}

// Fungsi rekursif - reverse string
func reverseString(s string) string {
	if len(s) <= 1 {
		return s
	}
	return string(s[len(s)-1]) + reverseString(s[:len(s)-1])
}

// Fungsi rekursif - binary search
func binarySearch(arr []int, target, left, right int) int {
	if left > right {
		return -1 // tidak ditemukan
	}
	
	mid := left + (right-left)/2
	
	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		return binarySearch(arr, target, left, mid-1)
	} else {
		return binarySearch(arr, target, mid+1, right)
	}
}

// Contoh penggunaan recursive functions
func DemoRecursiveFunctions() {
	fmt.Println("=== RECURSIVE FUNCTIONS ===")
	
	// Factorial
	fmt.Println("1. Factorial:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Factorial of %d: %d\n", i, factorial(i))
	}
	
	// Fibonacci
	fmt.Println("\n2. Fibonacci:")
	fmt.Print("Fibonacci sequence (first 10): ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()
	
	// Power
	fmt.Println("\n3. Power:")
	fmt.Printf("2^3 = %d\n", power(2, 3))
	fmt.Printf("5^4 = %d\n", power(5, 4))
	
	// GCD
	fmt.Println("\n4. Greatest Common Divisor:")
	fmt.Printf("GCD(48, 18) = %d\n", gcd(48, 18))
	fmt.Printf("GCD(100, 25) = %d\n", gcd(100, 25))
	
	// Sum Array
	fmt.Println("\n5. Sum Array (Recursive):")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Array: %v\n", numbers)
	fmt.Printf("Sum: %d\n", sumArray(numbers))
	
	// Reverse String
	fmt.Println("\n6. Reverse String:")
	original := "Hello"
	reversed := reverseString(original)
	fmt.Printf("Original: %s\n", original)
	fmt.Printf("Reversed: %s\n", reversed)
	
	// Binary Search
	fmt.Println("\n7. Binary Search:")
	sortedArray := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 7
	index := binarySearch(sortedArray, target, 0, len(sortedArray)-1)
	fmt.Printf("Array: %v\n", sortedArray)
	fmt.Printf("Searching for %d: ", target)
	if index != -1 {
		fmt.Printf("Found at index %d\n", index)
	} else {
		fmt.Println("Not found")
	}
	
	fmt.Println()
}