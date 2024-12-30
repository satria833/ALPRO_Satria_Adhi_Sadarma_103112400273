package main

import "fmt"

func main() {
        var totalBelanja int
        var mauKartu bool

        fmt.Print("Total belanja: Rp ")
        fmt.Scan(&totalBelanja)
        fmt.Print("Mau dibuatkan kartu? (true/false): ")
        fmt.Scan(&mauKartu)

        dapatDiskon := totalBelanja >= 100000

       
        dapatCashback := totalBelanja >= 200000 && mauKartu

        
        diskon := 0
        if dapatDiskon {
                diskon = totalBelanja * 10 / 100
        }

        
        hargaSetelahDiskon := totalBelanja - diskon

        cashback := 0
        if dapatCashback {
                cashback = 75000
        }

        
        hargaAkhir := hargaSetelahDiskon - cashback

       
        fmt.Println("Dapat kartu:", mauKartu)
        fmt.Println("Dapat diskon:", dapatDiskon)
        fmt.Println("Dapat cashback:", dapatCashback)
        fmt.Printf("Harga akhir: Rp %d\n", hargaAkhir)
}