package main

import "fmt"

func main() {
	var angka int

	fmt.Print("Masukan angka: ")
	fmt.Scan(&angka)

	if angka%2 == 0 {
		fmt.Println("angka adalah genap")
	} else {
		fmt.Println("angka adalah ganjil")
	}
}