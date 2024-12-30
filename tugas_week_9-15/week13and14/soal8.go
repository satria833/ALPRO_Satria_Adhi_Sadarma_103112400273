package main

import "fmt"

func main() {
        var x int
        fmt.Print("Masukkan bilangan positif: ")
        fmt.Scan(&x)

        for i := 1; i <= x; i++ {
                for j := 1; j <= i; j++ {
                        fmt.Print(j)
                }
                for j := i; j < x; j++ {
                        fmt.Print(j+1)
                }
                fmt.Println()
            }

        for i := x - 1; i >= 1; i-- {
                for j := 1; j <= i; j++ {
                        fmt.Print(j)
                }
                for j := i; j < x; j++ {
                        fmt.Print(j+1)
                }
                fmt.Println()
            }
    }