package main

import (
	"fmt"
	"sort"
	"strings"
)

// ========== FUNGSI DENGAN SLICE ==========

// Fungsi untuk mencari nilai maksimum dalam slice
func findMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

// Fungsi untuk mencari nilai minimum dalam slice
func findMin(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

// Fungsi untuk menghitung rata-rata
func average(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

// Fungsi untuk filter slice berdasarkan kondisi
func filter(numbers []int, condition func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		if condition(num) {
			result = append(result, num)
		}
	}
	return result
}

// Fungsi untuk menghapus duplikasi dalam slice
func removeDuplicates(slice []int) []int {
	keys := make(map[int]bool)
	var result []int
	
	for _, item := range slice {
		if !keys[item] {
			keys[item] = true
			result = append(result, item)
		}
	}
	return result
}

// ========== FUNGSI DENGAN MAP ==========

// Fungsi untuk menghitung kata dalam teks
func countWords(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}
	return wordCount
}

// Fungsi untuk menggabungkan dua map
func mergeMaps(map1, map2 map[string]int) map[string]int {
	result := make(map[string]int)
	
	// Copy map1
	for key, value := range map1 {
		result[key] = value
	}
	
	// Add map2 (akan menimpa jika key sama)
	for key, value := range map2 {
		result[key] += value // atau bisa += untuk menjumlahkan
	}
	
	return result
}

// Fungsi untuk mendapatkan keys dari map
func getKeys(m map[string]int) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// Fungsi untuk mendapatkan values dari map
func getValues(m map[string]int) []int {
	var values []int
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

// Fungsi untuk membalik key-value dalam map
func reverseMap(m map[string]int) map[int]string {
	result := make(map[int]string)
	for key, value := range m {
		result[value] = key
	}
	return result
}

// ========== FUNGSI UTILITAS UNTUK SLICE DAN MAP ==========

// Fungsi untuk mengecek apakah slice mengandung elemen tertentu
func contains(slice []int, item int) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}

// Fungsi untuk mengecek apakah map memiliki key tertentu
func hasKey(m map[string]int, key string) bool {
	_, exists := m[key]
	return exists
}

// Fungsi untuk generate slice angka
func generateRange(start, end int) []int {
	var result []int
	for i := start; i <= end; i++ {
		result = append(result, i)
	}
	return result
}

// Contoh penggunaan slice dan map functions
func DemoSliceMapFunctions() {
	fmt.Println("=== SLICE DAN MAP FUNCTIONS ===")
	
	// Slice operations
	fmt.Println("1. Slice Operations:")
	numbers := []int{3, 7, 2, 9, 1, 5, 7, 2, 9}
	fmt.Printf("Original slice: %v\n", numbers)
	fmt.Printf("Max: %d\n", findMax(numbers))
	fmt.Printf("Min: %d\n", findMin(numbers))
	fmt.Printf("Average: %.2f\n", average(numbers))
	
	// Filter slice
	evenNumbers := filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Even numbers: %v\n", evenNumbers)
	
	greaterThan5 := filter(numbers, func(n int) bool { return n > 5 })
	fmt.Printf("Numbers > 5: %v\n", greaterThan5)
	
	// Remove duplicates
	unique := removeDuplicates(numbers)
	fmt.Printf("Unique numbers: %v\n", unique)
	
	// Contains check
	fmt.Printf("Contains 7? %t\n", contains(numbers, 7))
	fmt.Printf("Contains 10? %t\n", contains(numbers, 10))
	
	// Generate range
	range1to10 := generateRange(1, 10)
	fmt.Printf("Range 1-10: %v\n", range1to10)
	
	// Map operations
	fmt.Println("\n2. Map Operations:")
	text := "hello world hello go world programming go"
	wordCount := countWords(text)
	fmt.Printf("Word count: %v\n", wordCount)
	
	// Get keys and values
	keys := getKeys(wordCount)
	values := getValues(wordCount)
	sort.Strings(keys) // Sort keys for consistent output
	fmt.Printf("Keys: %v\n", keys)
	fmt.Printf("Values: %v\n", values)
	
	// Check key existence
	fmt.Printf("Has key 'hello'? %t\n", hasKey(wordCount, "hello"))
	fmt.Printf("Has key 'python'? %t\n", hasKey(wordCount, "python"))
	
	// Merge maps
	map1 := map[string]int{"apple": 3, "banana": 2}
	map2 := map[string]int{"orange": 1, "apple": 1}
	merged := mergeMaps(map1, map2)
	fmt.Printf("Map1: %v\n", map1)
	fmt.Printf("Map2: %v\n", map2)
	fmt.Printf("Merged: %v\n", merged)
	
	// Reverse map
	reversed := reverseMap(map1)
	fmt.Printf("Original map: %v\n", map1)
	fmt.Printf("Reversed map: %v\n", reversed)
	
	fmt.Println()
}