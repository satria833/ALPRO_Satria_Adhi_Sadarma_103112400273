package main

import "fmt"

func containsDigit(n, x int) bool {
        for n > 0 {
                // Ambil digit terakhir dari n
                digit := n % 10

                // Jika digit yang ditemukan sama dengan x, maka return true
                if digit == x {
                        return true
                }

                // Buang digit terakhir
                n /= 10
        }

        // Jika tidak ditemukan digit yang sama, maka return false
        return false
}

func main() {
        var x, n int
        fmt.Print("Masukkan digit yang dicari (x): ")
        fmt.Scan(&x)
        fmt.Print("Masukkan bilangan (n): ")
        fmt.Scan(&n)

        if containsDigit(n, x) {
                fmt.Printf("true")
        } else {
                fmt.Printf("false")
        }
}