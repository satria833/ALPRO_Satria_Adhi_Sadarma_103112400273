package main

import "fmt"

func main() {
    var jumlahMahasiswa int

    fmt.Print("Masukkan jumlah mahasiswa: ")
    fmt.Scanln(&jumlahMahasiswa)

    jumlahMobil := jumlahMahasiswa / 7
	
    if jumlahMahasiswa % 7 != 0 {
        jumlahMobil++
    }

    fmt.Printf("Jumlah mobil yang dibutuhkan: %d\n", jumlahMobil)
}