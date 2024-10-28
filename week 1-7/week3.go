/*
# soal nomer 1
Sebuah program digunakan untuk mengecek suatu bilangan adalah ganjil atau bukan.
Masukan terdiri dari bilangan bulat positif.
Keluaran terdiri dari boolean yang menyatakan true apabila bilangan adalah ganjil, atau false apabila sebaliknya

# pseucode
algoritma

mulai
	input n
	Gunakan operator modulus 
	output true bilangan ganjil
	output false bilangan genap
end program
*/
package main

import "fmt"

func main() {
	var n int

	
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	isOdd := n%2 != 0

	fmt.Println(isOdd)
}
