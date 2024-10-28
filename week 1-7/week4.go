/*
#soal nomer 2
Sebuah program digunakan untuk menyimpan data persegi panjang yang berisi panjang, lebar, luas, dan keliling persegi panjang.
Masukan: terdiri dari dua bilangan riil yang menyatakan panjang dan lebar dari persegi panjang.
Keluaran: terdiri dari dua bilangan riil yang menyatakan luas dan keliling dari persegi panjang.
Catatan: Gunakan tipe bentukan untuk menyimpan data persegi panjang tersebut

#pseucode
algoritma

mulai 
	input p 
	input l
	luas p*l
	keliling 2*(p+l)
	output luas
	output keliling 
end progrmam
*/
package main

import "fmt"


type PersegiPanjang struct {
	panjang, lebar, luas, keliling float64
}


func (pp *PersegiPanjang) HitungLuasKeliling() {
	pp.luas = pp.panjang * pp.lebar
	pp.keliling = 2 * (pp.panjang + pp.lebar)
}

func main() {
	var panjang, lebar float64

	fmt.Print("Masukkan panjang persegi panjang: ")
	fmt.Scan(&panjang)
	fmt.Print("Masukkan lebar persegi panjang: ")
	fmt.Scan(&lebar)

	persegiPanjang := PersegiPanjang{panjang: panjang, lebar: lebar}

	persegiPanjang.HitungLuasKeliling()

	fmt.Printf("Luas persegi panjang: %.2f\n", persegiPanjang.luas)
	fmt.Printf("Keliling persegi panjang: %.2f\n", persegiPanjang.keliling)
}
