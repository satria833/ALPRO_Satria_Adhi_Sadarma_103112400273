package main

import "fmt"

func main() {
    var x int
    fmt.Print("Masukkan sebuah bilangan bulat positif: ")
    fmt.Scan(&x)

    for i := 1; i <= x; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print(i, " ")
        }
        fmt.Println()
    }
}