package main

import "fmt"

func main() {
        var n, lebar, lebarTerlebar int

       
        fmt.Print("Masukkan jumlah daun: ")
        fmt.Scan(&n)

       
        lebarTerlebar = 0

        
        for i := 0; i < n; i++ {
                fmt.Printf("Masukkan lebar daun ke-%d: ", i+1)
                fmt.Scan(&lebar)

                if lebar > lebarTerlebar {
                        lebarTerlebar = lebar
                }
        }

       
        fmt.Println("Lebar daun yang paling lebar adalah:", lebarTerlebar)
}