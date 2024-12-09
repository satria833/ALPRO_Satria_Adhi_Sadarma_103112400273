package main

import (
	"fmt"
)

func main() {
	var totalBelanja float64
	var namaBarang string
	var hargaBarang float64
	var lanjutkanTransaksi string

	fmt.Println("=== Aplikasi Kasir Sederhana ===")
	for {
		fmt.Print("Masukkan nama barang: ")
		fmt.Scan(&namaBarang)

		fmt.Print("Masukkan harga barang: Rp ")
		fmt.Scan(&hargaBarang)
	
		totalBelanja += hargaBarang
		
		fmt.Printf("Barang: %s, Harga: Rp %.2f\n", namaBarang, hargaBarang)
		fmt.Printf("Total belanja sementara: Rp %.2f\n", totalBelanja)
		
		fmt.Print("Apakah Anda ingin menambahkan barang lain? (ya/tidak): ")
		fmt.Scan(&lanjutkanTransaksi)
		
		if lanjutkanTransaksi == "tidak" {
			break
		}
	}	
	fmt.Printf("\nTotal belanja Anda: Rp %.2f\n", totalBelanja)
	fmt.Println("Terima kasih, selamat datang kembali!")
}
