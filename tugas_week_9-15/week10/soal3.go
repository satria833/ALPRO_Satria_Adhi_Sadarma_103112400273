package main

import "fmt"

func main() {
        var keuntunganBulanIni, keuntunganBulanLalu float64

        fmt.Print("Masukkan keuntungan bulan ini: ")
        fmt.Scanln(&keuntunganBulanIni)

        fmt.Print("Masukkan keuntungan bulan lalu: ")
        fmt.Scanln(&keuntunganBulanLalu)

        perubahanKeuntungan := keuntunganBulanIni - keuntunganBulanLalu

        if perubahanKeuntungan > 0 {
                fmt.Printf("Keuntungan meningkat sebesar %.2f\n", perubahanKeuntungan)
        } else if perubahanKeuntungan < 0 {
                fmt.Printf("Keuntungan menurun sebesar %.2f\n", -perubahanKeuntungan)
        } else {
                fmt.Println("tetap")
        }
}