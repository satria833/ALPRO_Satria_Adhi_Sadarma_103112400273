package main

import "fmt"

func main() {
    var bilangan int

    fmt.Print("Masukkan bilangan: ")
    fmt.Scanln(&bilangan)

    if bilangan%2 == 0 && bilangan < 0 {
        fmt.Println("Bilangan adalah genap negatif")
    } else {
        fmt.Println("Bilangan bukan genap negatif")
    }
}