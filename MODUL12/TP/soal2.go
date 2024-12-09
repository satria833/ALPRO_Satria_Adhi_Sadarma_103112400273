package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Println("Selamat datang! Masukkan kata apa saja.")
	fmt.Println("Ketik 'telkom' untuk menyelesaikan program.")

	for {
		// Meminta input dari pengguna
		fmt.Print("Masukkan kata: ")
		fmt.Scan(&input)

		if strings.ToLower(input) == "telkom" {
			fmt.Println("Program selesai.")
			break 
		} else {
			fmt.Println("Kata diterima. Silakan masukkan kata lainnya.")
		}
	}
}
