package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64

	
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&radius)

	
	area := math.Pi * radius * radius           // Luas = π * r^2
	circumference := 2 * math.Pi * radius       // Keliling = 2 * π * r

	
	fmt.Printf("Luas lingkaran: %.2f\n", area)
	fmt.Printf("Keliling lingkaran: %.2f\n", circumference)
}
