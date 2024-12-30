package main

import "fmt"

func main() {
    var x, y int

    fmt.Print("Masukkan bilangan pertama (x): ")
    fmt.Scanln(&x)

    fmt.Print("Masukkan bilangan kedua (y): ")
    fmt.Scanln(&y)

    xAdalahFaktorDariY := y%x == 0

    yAdalahFaktorDariX := x%y == 0

    fmt.Println("x adalah faktor dari y:", xAdalahFaktorDariY)
    fmt.Println("y adalah faktor dari x:", yAdalahFaktorDariX)
}