package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("masukan nilai: ")
	fmt.Scan(&nilai)

	if nilai >= 90 {
		fmt.Println("A")
	} else if nilai >= 80 && nilai < 90 {
		fmt.Println("AB")
	} else if nilai >= 70 && nilai < 80 {
		fmt.Println("B")
	} else if nilai <= 70 {
		fmt.Println("C")
	}
}