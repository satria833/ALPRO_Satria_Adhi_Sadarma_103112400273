package main

import "fmt"

func main() {
    var sisiA, sisiB, sisiC int

    fmt.Print("Masukkan panjang sisi A: ")
    fmt.Scanln(&sisiA)

    fmt.Print("Masukkan panjang sisi B: ")
    fmt.Scanln(&sisiB)

    fmt.Print("Masukkan panjang sisi C: ")
    fmt.Scanln(&sisiC)

    if sisiA == sisiB && sisiB == sisiC {
        fmt.Println("Segitiga Sama Sisi")
    } else if sisiA == sisiB || sisiA == sisiC || sisiB == sisiC {
        fmt.Println("Segitiga Sama Kaki")
    } else {
        fmt.Println("Segitiga Sembarang")
    }
}