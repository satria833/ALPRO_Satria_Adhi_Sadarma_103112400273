package main

import "fmt"

func main () {
//input jumblah barang
	var poin,barang int 
	fmt.Print("masukan jumlah barang yang dibeli: ")
	fmt.Scan(&barang)
	
	// menghitung sebuah poin dengan sebuah perulangan
	for i := 1; i <= barang; i++ {
		if i <= 5 {
			poin = poin + 10 // 5 barang pertama memberi 10 poin per barang
		} else {
			poin = poin + 15 // Barang setelah barang kelima memberi 15 poin per barang
		}
	}
fmt.Println("total poin yang didapat: ", poin, "poin") // menampilkan hasil
}l