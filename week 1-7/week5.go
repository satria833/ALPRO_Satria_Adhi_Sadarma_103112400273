/*
soal nomer 3
Diberikan sebuah daftar bilangan, kalikan setiap bilangan dengan indeksnya (indeks dimulai dari 1 hingga ke n). Buatlah program  untuk menghitung hasil dari perkalian ini untuk setiap bilangan.
Masukan: Baris pertama adalah angka n yang menunjukkan jumlah bilangan, diikuti oleh n bilangan.
Keluaran: Keluaran adalah hasil perkalian setiap bilangan dengan indeksnya.

# pseucode
algoritma

mulai

	input n
	perulangan dari 0 hingga n-1
	output hasil perkalian

end program
*/
package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan jumlah bilangan: ")
	fmt.Scan(&n)

	numbers := make([]int, n)
	fmt.Println("Masukkan bilangan-bilangannya:")
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	fmt.Println("Hasil perkalian setiap bilangan dengan indeksnya:")
	for i, num := range numbers {
		hasil := num * (i + 1)
		fmt.Printf("Bilangan %d dikali indeks %d = %d\n", num, i+1, hasil)
	}
}
