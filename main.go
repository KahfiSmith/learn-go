package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=== LEARN GO - FUNGSI-FUNGSI GO ===")
	fmt.Println("Pilih kategori fungsi yang ingin dipelajari:")
	fmt.Println()
	fmt.Println("1. Basic Functions (Fungsi Dasar)")
	fmt.Println("2. Advanced Functions (Higher-order, Closure)")
	fmt.Println("3. Recursive Functions (Fungsi Rekursif)")
	fmt.Println("4. Struct & Methods (Struct dan Interface)")
	fmt.Println("5. Slice & Map Functions (Operasi Slice dan Map)")
	fmt.Println("6. Concurrency Functions (Goroutine & Channel)")
	fmt.Println("7. Error Handling (Defer, Panic, Recover)")
	fmt.Println("8. Utility Functions (Fungsi Utilitas)")
	fmt.Println("9. Jalankan Semua Contoh")
	fmt.Println("0. Keluar")
	fmt.Println()
	
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("Masukkan pilihan (0-9): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid. Masukkan angka 0-9.")
			continue
		}
		
		fmt.Println()
		
		switch choice {
		case 1:
			DemoBasicFunctions()
		case 2:
			DemoAdvancedFunctions()
		case 3:
			DemoRecursiveFunctions()
		case 4:
			DemoStructMethods()
		case 5:
			DemoSliceMapFunctions()
		case 6:
			DemoConcurrencyFunctions()
		case 7:
			DemoErrorHandling()
		case 8:
			DemoUtilityFunctions()
		case 9:
			fmt.Println("=== MENJALANKAN SEMUA CONTOH ===\n")
			DemoBasicFunctions()
			DemoAdvancedFunctions()
			DemoRecursiveFunctions()
			DemoStructMethods()
			DemoSliceMapFunctions()
			DemoConcurrencyFunctions()
			DemoErrorHandling()
			DemoUtilityFunctions()
			fmt.Println("=== SEMUA CONTOH SELESAI ===\n")
		case 0:
			fmt.Println("Terima kasih! Selamat belajar Go!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Masukkan angka 0-9.")
			continue
		}
		
		fmt.Println("Tekan Enter untuk kembali ke menu...")
		reader.ReadString('\n')
		fmt.Println()
		fmt.Println("=== LEARN GO - FUNGSI-FUNGSI GO ===")
		fmt.Println("Pilih kategori fungsi yang ingin dipelajari:")
		fmt.Println()
		fmt.Println("1. Basic Functions (Fungsi Dasar)")
		fmt.Println("2. Advanced Functions (Higher-order, Closure)")
		fmt.Println("3. Recursive Functions (Fungsi Rekursif)")
		fmt.Println("4. Struct & Methods (Struct dan Interface)")
		fmt.Println("5. Slice & Map Functions (Operasi Slice dan Map)")
		fmt.Println("6. Concurrency Functions (Goroutine & Channel)")
		fmt.Println("7. Error Handling (Defer, Panic, Recover)")
		fmt.Println("8. Utility Functions (Fungsi Utilitas)")
		fmt.Println("9. Jalankan Semua Contoh")
		fmt.Println("0. Keluar")
		fmt.Println()
	}
}