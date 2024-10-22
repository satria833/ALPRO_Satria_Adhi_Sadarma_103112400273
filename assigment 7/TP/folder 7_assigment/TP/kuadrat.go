package main

import "fmt"

func main() {
	// input masukan
	var n int
	fmt.Print("Masukan jumlah n kuadrat: ")
	fmt.Scan(&n)

	//mencetak hasil kuadrat dari 1 sampai ke n
	for i := 1; i <= n; i++ {
		if i == n {
			fmt.Printf("%d", i*i)
		} else {
			fmt.Printf("%d", i*i)
			fmt.Print(" ") // mencetak hasil dengan ditambah spasi
		}
	}
}