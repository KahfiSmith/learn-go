package main

import (
	"fmt"
	"math"
)

// ========== STRUCT DAN METHODS ==========

// Person struct
type Person struct {
	Name string
	Age  int
	Email string
}

// Method dengan value receiver
func (p Person) GetInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d, Email: %s", p.Name, p.Age, p.Email)
}

// Method dengan pointer receiver
func (p *Person) SetAge(age int) {
	p.Age = age
}

func (p *Person) SetEmail(email string) {
	p.Email = email
}

// Method untuk validasi
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// ========== INTERFACE ==========

// Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
	GetType() string
}

// Rectangle struct
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) GetType() string {
	return "Rectangle"
}

// Circle struct
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) GetType() string {
	return "Circle"
}

// Triangle struct
type Triangle struct {
	Base, Height, Side1, Side2 float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return t.Base + t.Side1 + t.Side2
}

func (t Triangle) GetType() string {
	return "Triangle"
}

// ========== FUNGSI YANG BEKERJA DENGAN INTERFACE ==========

// Fungsi yang menerima interface
func printShapeInfo(s Shape) {
	fmt.Printf("%s - Area: %.2f, Perimeter: %.2f\n", 
		s.GetType(), s.Area(), s.Perimeter())
}

// Fungsi untuk menghitung total area dari slice shapes
func calculateTotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

// ========== EMBEDDED STRUCT ==========

// Address struct
type Address struct {
	Street  string
	City    string
	ZipCode string
}

// Employee struct dengan embedded Address
type Employee struct {
	Person  // embedded struct
	Address // embedded struct
	Salary  float64
	JobTitle string
}

// Method untuk Employee
func (e Employee) GetFullInfo() string {
	return fmt.Sprintf("%s works as %s, lives at %s, %s, earns $%.2f",
		e.Name, e.JobTitle, e.Street, e.City, e.Salary)
}

// Contoh penggunaan struct dan methods
func DemoStructMethods() {
	fmt.Println("=== STRUCT DAN METHODS ===")
	
	// Basic struct dan methods
	fmt.Println("1. Basic Struct dan Methods:")
	person := Person{Name: "Alice", Age: 25, Email: "alice@example.com"}
	fmt.Println(person.GetInfo())
	fmt.Printf("Is adult? %t\n", person.IsAdult())
	
	// Pointer receiver methods
	person.SetAge(26)
	person.SetEmail("alice.new@example.com")
	fmt.Println("After update:", person.GetInfo())
	
	// Interface
	fmt.Println("\n2. Interface:")
	shapes := []Shape{
		Rectangle{Width: 5, Height: 3},
		Circle{Radius: 4},
		Triangle{Base: 6, Height: 4, Side1: 5, Side2: 5},
	}
	
	for _, shape := range shapes {
		printShapeInfo(shape)
	}
	
	totalArea := calculateTotalArea(shapes)
	fmt.Printf("Total area of all shapes: %.2f\n", totalArea)
	
	// Embedded struct
	fmt.Println("\n3. Embedded Struct:")
	employee := Employee{
		Person: Person{Name: "Bob", Age: 30, Email: "bob@company.com"},
		Address: Address{Street: "123 Main St", City: "New York", ZipCode: "10001"},
		Salary: 75000,
		JobTitle: "Software Engineer",
	}
	
	fmt.Println(employee.GetFullInfo())
	// Dapat mengakses field embedded langsung
	fmt.Printf("Employee name: %s\n", employee.Name)
	fmt.Printf("Employee city: %s\n", employee.City)
	
	fmt.Println()
}