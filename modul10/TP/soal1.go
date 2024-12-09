package main

import (
	"fmt"
)

func main() {
	
	fmt.Println("=== Menu Restoran Cepat Saji ===")
	fmt.Println("1. Burger       - Rp 25,000")
	fmt.Println("2. Ayam Goreng  - Rp 20,000")
	fmt.Println("3. Kentang Goreng - Rp 15,000")
	fmt.Println("4. Es Teh       - Rp 8,000")
	fmt.Println("5. Jus Jeruk    - Rp 10,000")
	fmt.Println("===============================")

	
	var kodeItem int
	fmt.Print("Masukkan kode menu (1-5): ")
	fmt.Scan(&kodeItem)

	switch kodeItem {
	case 1:
		fmt.Println("Anda memilih: Burger - Rp 25,000")
	case 2:
		fmt.Println("Anda memilih: Ayam Goreng - Rp 20,000")
	case 3:
		fmt.Println("Anda memilih: Kentang Goreng - Rp 15,000")
	case 4:
		fmt.Println("Anda memilih: Es Teh - Rp 8,000")
	case 5:
		fmt.Println("Anda memilih: Jus Jeruk - Rp 10,000")
	default:
		fmt.Println("Kode menu tidak valid.")
	}
}
