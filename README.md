# Learn Go - Fungsi-Fungsi Go

Proyek ini berisi contoh-contoh lengkap fungsi-fungsi dalam bahasa pemrograman Go, disusun secara terorganisir dalam file-file terpisah berdasarkan kategori.

## Struktur File

### 1. `main.go` - Program Utama
File utama yang berisi menu interaktif untuk menjalankan contoh-contoh fungsi berdasarkan kategori.

### 2. `basic_functions.go` - Fungsi Dasar
Berisi contoh fungsi-fungsi dasar Go:
- Fungsi sederhana tanpa parameter dan return value
- Fungsi dengan parameter
- Fungsi dengan return value
- Fungsi dengan multiple return values
- Fungsi dengan named return values
- Fungsi dengan variadic parameters (parameter tak terbatas)

**Contoh:**
```go
func add(a, b int) int {
    return a + b
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("tidak bisa dibagi dengan nol")
    }
    return a / b, nil
}
```

### 3. `advanced_functions.go` - Fungsi Lanjutan
Berisi contoh fungsi-fungsi lanjutan:
- Function as first-class citizens
- Higher-order functions
- Closure
- Function yang mengembalikan function
- Function validation patterns

**Contoh:**
```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

### 4. `recursive_functions.go` - Fungsi Rekursif
Berisi contoh fungsi-fungsi rekursif:
- Factorial
- Fibonacci
- Power calculation
- Greatest Common Divisor (GCD)
- Sum array dengan rekursi
- Reverse string dengan rekursi
- Binary search

**Contoh:**
```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}
```

### 5. `struct_methods.go` - Struct dan Methods
Berisi contoh penggunaan struct dan methods:
- Value receiver methods
- Pointer receiver methods
- Interface implementation
- Embedded struct
- Polymorphism dengan interface

**Contoh:**
```go
type Person struct {
    Name string
    Age  int
}

func (p Person) GetInfo() string {
    return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func (p *Person) SetAge(age int) {
    p.Age = age
}
```

### 6. `slice_map_functions.go` - Operasi Slice dan Map
Berisi fungsi-fungsi untuk bekerja dengan slice dan map:
- Operasi pencarian (max, min, average)
- Filter slice dengan kondisi
- Remove duplicates
- Word counting dengan map
- Merge maps
- Get keys/values dari map

**Contoh:**
```go
func filter(numbers []int, condition func(int) bool) []int {
    var result []int
    for _, num := range numbers {
        if condition(num) {
            result = append(result, num)
        }
    }
    return result
}
```

### 7. `concurrency_functions.go` - Goroutine dan Channel
Berisi contoh penggunaan concurrency:
- Worker pool pattern
- Channel communication
- WaitGroup untuk synchronization
- Mutex untuk thread safety
- Select statement
- Fan-in/Fan-out patterns
- Buffered channels

**Contoh:**
```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        results <- job * 2
    }
}
```

### 8. `error_handling.go` - Error Handling
Berisi contoh error handling di Go:
- Defer statement
- Panic dan recover
- Resource management
- Multiple defer (LIFO order)
- Real-world error handling patterns

**Contoh:**
```go
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
```

### 9. `utility_functions.go` - Fungsi Utilitas
Berisi fungsi-fungsi utilitas yang berguna:
- String manipulation (validation, cleaning, formatting)
- Math utilities (temperature conversion, distance calculation)
- Number utilities (prime numbers, formatting)
- Date/time utilities
- Validation functions
- Conversion utilities

**Contoh:**
```go
func celsiusToFahrenheit(celsius float64) float64 {
    return (celsius * 9 / 5) + 32
}

func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}
```

## Cara Menjalankan

1. **Jalankan program utama:**
   ```bash
   go run .
   ```

2. **Pilih kategori yang ingin dipelajari:**
   - Program akan menampilkan menu dengan 9 pilihan kategori
   - Masukkan angka 1-9 untuk melihat contoh dari kategori tertentu
   - Pilih 9 untuk menjalankan semua contoh sekaligus
   - Pilih 0 untuk keluar

3. **Jalankan file tertentu:**
   ```bash
   go run main.go basic_functions.go
   go run main.go advanced_functions.go
   # dst...
   ```

## Konsep Yang Dipelajari

### Fungsi Dasar
- Parameter dan return values
- Variadic functions
- Named returns

### Fungsi Lanjutan
- First-class functions
- Higher-order functions
- Closures
- Function composition

### Rekursi
- Base case dan recursive case
- Tail recursion
- Aplikasi rekursi dalam algoritma

### OOP dengan Go
- Struct dan methods
- Interface dan polymorphism
- Embedding

### Concurrency
- Goroutines
- Channels
- Synchronization
- Patterns untuk concurrent programming

### Error Handling
- Error sebagai value
- Panic dan recover
- Resource management dengan defer

### Best Practices
- Code organization
- Error handling patterns
- Performance considerations
- Memory management

## Tips Belajar

1. **Mulai dari basic:** Pelajari fungsi dasar terlebih dahulu sebelum ke topik lanjutan
2. **Practice:** Coba modifikasi contoh-contoh yang ada
3. **Experiment:** Buat fungsi sendiri berdasarkan pola yang dipelajari
4. **Read documentation:** Baca dokumentasi Go untuk memahami lebih dalam
5. **Build projects:** Aplikasikan konsep dalam proyek nyata

## Requirements

- Go 1.16 atau lebih baru
- Terminal/Command prompt untuk menjalankan program

## Kontribusi

Jika ingin menambah contoh fungsi atau memperbaiki yang sudah ada, silakan buat pull request atau issue.

---

**Happy Learning Go! 🚀**