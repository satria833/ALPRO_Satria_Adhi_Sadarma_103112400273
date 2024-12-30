package main

import (
	"fmt"
)

func main() {
	var p, r, hari int

	
	fmt.Print("Masukkan jumlah halaman novel: ")
	fmt.Scan(&p)
	fmt.Print("Masukkan jumlah halaman yang dibaca per hari: ")
	fmt.Scan(&r)

	hari = p / r
	if p%r > 0 {
		hari += 1
	}

	
	fmt.Printf("%d hari\n", hari)
}