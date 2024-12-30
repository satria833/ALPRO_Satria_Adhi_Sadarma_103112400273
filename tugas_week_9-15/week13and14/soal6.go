package main

import (
        "fmt"
        
)

func main() {
        var input string
        var countA, countB, countC int

        fmt.Println("Masukkan data kendaraan (akhiri dengan karakter selain A, B, atau C):")

        for {
                fmt.Scan(&input)

               
                if input != "A" && input != "B" && input != "C" {
                        break
                }

                
                switch input {
                case "A":
                        countA++
                case "B":
                        countB++
                case "C":
                        countC++
                }
        }

        fmt.Println("Tipe A =", countA)
        fmt.Println("Tipe B =", countB)
        fmt.Println("Tipe C =", countC)
}