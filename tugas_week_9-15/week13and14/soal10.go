package main

import "fmt"

func main() {
    var x int
    fmt.Print("Masukkan sebuah bilangan bulat positif: ")
    fmt.Scan(&x)

    for i := 1; i <= x; i++ {
        for j := 1; j <= x; j++ {
            if j == i || j == (x-i+1) {
                fmt.Print(i, " ")
            } else {
                fmt.Print("  ")
            }
        }
        fmt.Println()
    }
}