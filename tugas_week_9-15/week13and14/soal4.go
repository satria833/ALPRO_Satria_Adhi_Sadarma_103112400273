package main

import "fmt"

func digitTerbesar(angka int) int {
        digitTerbesar := 0

        for angka > 0 {
                
                digit := angka % 10

               
                if digit > digitTerbesar {
                        digitTerbesar = digit
                }

               
                angka /= 10
        }

        return digitTerbesar
}

func main() {
        var bilangan int
        fmt.Print("Masukkan bilangan: ")
        fmt.Scan(&bilangan)

        hasil := digitTerbesar(bilangan)
        fmt.Println("Digit terbesar adalah:", hasil)
}