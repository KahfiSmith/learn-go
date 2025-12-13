package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ========== STRING UTILITIES ==========

// Fungsi untuk validasi email sederhana
func isValidEmail(email string) bool {
	// Simple email validation
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// Fungsi untuk validasi email dengan regex
func isValidEmailRegex(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

// Fungsi untuk membersihkan string
func cleanString(s string) string {
	// Remove extra spaces and trim
	cleaned := strings.TrimSpace(s)
	// Replace multiple spaces with single space
	re := regexp.MustCompile(`\s+`)
	cleaned = re.ReplaceAllString(cleaned, " ")
	return cleaned
}

// Fungsi untuk capitalize first letter of each word
func toTitleCase(s string) string {
	return strings.Title(strings.ToLower(s))
}

// Fungsi untuk reverse string
func reverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Fungsi untuk check palindrome
func isPalindrome(s string) bool {
	// Clean string: remove spaces and convert to lowercase
	cleaned := strings.ReplaceAll(strings.ToLower(s), " ", "")
	return cleaned == reverseStr(cleaned)
}

// ========== MATH UTILITIES ==========

// Fungsi untuk konversi temperature
func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

// Fungsi untuk menghitung jarak antara dua titik
func distance(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}

// Fungsi untuk menghitung area dan volume
func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func sphereVolume(radius float64) float64 {
	return (4.0 / 3.0) * math.Pi * radius * radius * radius
}

func rectangleArea(width, height float64) float64 {
	return width * height
}

func cylinderVolume(radius, height float64) float64 {
	return math.Pi * radius * radius * height
}

// ========== NUMBER UTILITIES ==========

// Fungsi untuk check if number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Fungsi untuk generate prime numbers up to n
func generatePrimes(n int) []int {
	var primes []int
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

// Fungsi untuk format number dengan separator
func formatNumber(n int) string {
	str := strconv.Itoa(n)
	if len(str) <= 3 {
		return str
	}
	
	var result []string
	for i := len(str); i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}
		result = append([]string{str[start:i]}, result...)
	}
	return strings.Join(result, ",")
}

// Fungsi untuk round to specific decimal places
func roundToDecimal(num float64, places int) float64 {
	multiplier := math.Pow(10, float64(places))
	return math.Round(num*multiplier) / multiplier
}

// ========== DATE AND TIME UTILITIES ==========

// Fungsi untuk format date
func formatDate(t time.Time, format string) string {
	switch format {
	case "DD/MM/YYYY":
		return t.Format("02/01/2006")
	case "MM/DD/YYYY":
		return t.Format("01/02/2006")
	case "YYYY-MM-DD":
		return t.Format("2006-01-02")
	case "DD Mon YYYY":
		return t.Format("02 Jan 2006")
	default:
		return t.Format(time.RFC3339)
	}
}

// Fungsi untuk menghitung umur
func calculateAge(birthDate time.Time) int {
	now := time.Now()
	age := now.Year() - birthDate.Year()
	
	if now.YearDay() < birthDate.YearDay() {
		age--
	}
	
	return age
}

// Fungsi untuk menghitung selisih waktu
func timeDifference(start, end time.Time) string {
	diff := end.Sub(start)
	
	days := int(diff.Hours() / 24)
	hours := int(diff.Hours()) % 24
	minutes := int(diff.Minutes()) % 60
	seconds := int(diff.Seconds()) % 60
	
	return fmt.Sprintf("%d days, %d hours, %d minutes, %d seconds", days, hours, minutes, seconds)
}

// Fungsi untuk check if year is leap year
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

// ========== VALIDATION UTILITIES ==========

// Fungsi untuk validasi nomor telepon Indonesia
func isValidPhoneNumber(phone string) bool {
	// Remove all non-digit characters
	cleaned := regexp.MustCompile(`[^\d]`).ReplaceAllString(phone, "")
	
	// Check if starts with +62 or 08 and has proper length
	if strings.HasPrefix(cleaned, "62") && len(cleaned) >= 11 && len(cleaned) <= 13 {
		return true
	}
	if strings.HasPrefix(cleaned, "08") && len(cleaned) >= 10 && len(cleaned) <= 12 {
		return true
	}
	return false
}

// Fungsi untuk validasi NIK (Nomor Induk Kependudukan)
func isValidNIK(nik string) bool {
	// NIK should be 16 digits
	if len(nik) != 16 {
		return false
	}
	
	// Check if all characters are digits
	for _, char := range nik {
		if char < '0' || char > '9' {
			return false
		}
	}
	
	return true
}

// ========== CONVERSION UTILITIES ==========

// Fungsi untuk convert bytes to human readable format
func bytesToHuman(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	
	units := []string{"KB", "MB", "GB", "TB", "PB"}
	return fmt.Sprintf("%.1f %s", float64(bytes)/float64(div), units[exp])
}

// Fungsi untuk generate random string
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(result)
}

// Contoh penggunaan utility functions
func DemoUtilityFunctions() {
	fmt.Println("=== UTILITY FUNCTIONS ===")
	
	// String utilities
	fmt.Println("1. String Utilities:")
	email := "user@example.com"
	fmt.Printf("Email '%s' valid? %t\n", email, isValidEmail(email))
	fmt.Printf("Email '%s' valid (regex)? %t\n", email, isValidEmailRegex(email))
	
	messy := "   hello    world   go   "
	fmt.Printf("Original: '%s'\n", messy)
	fmt.Printf("Cleaned: '%s'\n", cleanString(messy))
	
	name := "john doe smith"
	fmt.Printf("Title case: %s\n", toTitleCase(name))
	
	word := "racecar"
	fmt.Printf("'%s' is palindrome? %t\n", word, isPalindrome(word))
	
	// Math utilities
	fmt.Println("\n2. Math Utilities:")
	celsius := 25.0
	fmt.Printf("%.1f°C = %.1f°F\n", celsius, celsiusToFahrenheit(celsius))
	fmt.Printf("%.1f°C = %.1f K\n", celsius, celsiusToKelvin(celsius))
	
	radius := 5.0
	fmt.Printf("Circle area (r=%.1f): %.2f\n", radius, circleArea(radius))
	fmt.Printf("Sphere volume (r=%.1f): %.2f\n", radius, sphereVolume(radius))
	
	dist := distance(0, 0, 3, 4)
	fmt.Printf("Distance from (0,0) to (3,4): %.2f\n", dist)
	
	// Number utilities
	fmt.Println("\n3. Number Utilities:")
	num := 17
	fmt.Printf("%d is prime? %t\n", num, isPrime(num))
	
	primes := generatePrimes(20)
	fmt.Printf("Primes up to 20: %v\n", primes)
	
	bigNumber := 1234567
	fmt.Printf("Formatted number: %s\n", formatNumber(bigNumber))
	
	decimal := 3.14159265
	fmt.Printf("Rounded to 2 places: %.2f\n", roundToDecimal(decimal, 2))
	
	// Date/Time utilities
	fmt.Println("\n4. Date/Time Utilities:")
	now := time.Now()
	fmt.Printf("Current time (DD/MM/YYYY): %s\n", formatDate(now, "DD/MM/YYYY"))
	fmt.Printf("Current time (DD Mon YYYY): %s\n", formatDate(now, "DD Mon YYYY"))
	
	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
	age := calculateAge(birthDate)
	fmt.Printf("Age for someone born on %s: %d years\n", formatDate(birthDate, "DD/MM/YYYY"), age)
	
	fmt.Printf("Year 2024 is leap year? %t\n", isLeapYear(2024))
	
	// Validation utilities
	fmt.Println("\n5. Validation Utilities:")
	phones := []string{"081234567890", "+6281234567890", "021-12345678"}
	for _, phone := range phones {
		fmt.Printf("Phone '%s' valid? %t\n", phone, isValidPhoneNumber(phone))
	}
	
	nik := "1234567890123456"
	fmt.Printf("NIK '%s' valid? %t\n", nik, isValidNIK(nik))
	
	// Conversion utilities
	fmt.Println("\n6. Conversion Utilities:")
	bytes := int64(1024*1024*500 + 1024*256) // 500.25 MB
	fmt.Printf("%d bytes = %s\n", bytes, bytesToHuman(bytes))
	
	randomStr := generateRandomString(8)
	fmt.Printf("Random string (length 8): %s\n", randomStr)
	
	fmt.Println()
}