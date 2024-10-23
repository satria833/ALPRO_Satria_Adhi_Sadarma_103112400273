package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&N)

	if N > 0 {
		for i := 1; i <= N; i++ {
			fmt.Println(i * i)
		}
	} else {
		fmt.Println("Nilai N harus bilangan bulat positif.")
	}
}
