package main

import (
	"fmt"
)

func main() {
	
	const angkaRahasia = 7
	var tebak int

	fmt.Println("Selamat datang di permainan tebak angka!")
	fmt.Println("Tebak angka rahasia antara 1 hingga 10.")

	for {
		
		fmt.Print("Masukkan tebakan Anda: ")
		fmt.Scan(&tebak)

		if tebak == angkaRahasia {
			fmt.Println("Selamat! Tebakan Anda benar.")
			break 
		} else {
			fmt.Println("Tebakan Anda salah. Coba lagi.")
		}
	}

	fmt.Println("Terima kasih telah bermain!")
}
