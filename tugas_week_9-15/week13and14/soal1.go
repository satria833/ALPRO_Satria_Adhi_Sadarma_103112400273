package main

import "fmt"

func isPrime(num int) bool {
        // Bilangan prima lebih besar dari 1
        if num <= 1 {
                return false
        }

        // Cek kelipatan mulai dari 2 sampai akar kuadrat dari num
        for i := 2; i*i <= num; i++ {
                if num%i == 0 {
                        return false
                }
        }

        return true
}

func main() {
        var x int
        fmt.Print("Masukkan bilangan: ")
        fmt.Scan(&x)

        if isPrime(x) {
                fmt.Println(x, "true")
        } else {
                fmt.Println(x, "false")
        }
}