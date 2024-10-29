package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Masukan nilai ujian:")
	fmt.Scan(&nilai)

	if nilai >= 70 {
		fmt.Println("lulus")
	} else {
		fmt.Println("tidak lulus")
	}
}