package main

import (
        "fmt"
)

func main() {
        var suhu1, suhu2, suhu3, suhu4, suhu5 float64

        fmt.Print("Masukkan suhu pertama: ")
        fmt.Scanln(&suhu1)
        fmt.Print("Masukkan suhu kedua: ")
        fmt.Scanln(&suhu2)
        fmt.Print("Masukkan suhu ketiga: ")
        fmt.Scanln(&suhu3)
        fmt.Print("Masukkan suhu keempat: ")
        fmt.Scanln(&suhu4)
        fmt.Print("Masukkan suhu kelima: ")
        fmt.Scanln(&suhu5)

        kenaikan := suhu2 > suhu1 && suhu3 > suhu2 && suhu4 > suhu3 && suhu5 > suhu4
        penurunan := suhu2 < suhu1 && suhu3 < suhu2 && suhu4 < suhu3 && suhu5 < suhu4

        if kenaikan {
                fmt.Println("Stabil naik")
        } else if penurunan {
                fmt.Println("Stabil turun")
        } else {
                fmt.Println("Tidak stabil")
        }
}