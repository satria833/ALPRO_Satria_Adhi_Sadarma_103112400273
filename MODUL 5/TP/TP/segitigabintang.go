package main

import "fmt"

func main() {
	var n int

	fmt.Println("Masukan bilangan yang ingin di loop: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ { 
		fmt.Println("*") 
	}
}