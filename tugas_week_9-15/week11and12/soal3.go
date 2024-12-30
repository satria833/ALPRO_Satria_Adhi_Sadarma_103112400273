package main

import "fmt"

func main() {
        var bilangan int

        fmt.Print("Masukkan bilangan bulat positif: ")
        fmt.Scan(&bilangan)

        var jumlahDigit, jumlahTotal int
        var digit int

        for bilangan > 0 {
              
                digit = bilangan % 10

                
                fmt.Print(digit, " ")

               
                jumlahTotal += digit

               
                bilangan /= 10
                jumlahDigit++
        }

        fmt.Println()
        fmt.Printf("Jumlah digit: %d\n", jumlahDigit)
        fmt.Printf("Jumlah total digit: %d\n", jumlahTotal)
}