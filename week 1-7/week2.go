/*
soal nomer 2
Sebuah program digunakan untuk menghitung luas dan keliling lingkaran.
Luas lingkaran = πr² Keliling lingkaran = 2πr
Masukan terdiri dari suatu bilangan riil yang menyatakan jari-jari lingkaran.
Keluaran terdiri dari dua bilangan yang menyatakan luas dan keliling lingkaran

# pseucode
algoritma

mulai
	input r
	luas π * r * r
	output Luas lingkaran [nilai luas]
	output Keliling lingkaran [nilai keliling]
end program


*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64

	
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&r)

	luas := math.Pi * r * r
	keliling := 2 * math.Pi * r

	fmt.Printf("Luas lingkaran: %.2f\n", luas)
	fmt.Printf("Keliling lingkaran: %.2f\n", keliling)
}
