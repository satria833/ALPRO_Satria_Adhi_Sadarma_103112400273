package main

import "fmt"

func main() {
        var saldo, transaksi int

        fmt.Println("Masukkan transaksi (angka negatif untuk pengeluaran, 0 untuk berhenti):")

        for {
                fmt.Scan(&transaksi)

                if transaksi == 0 {
                        break
                }

                saldo += transaksi
        }

        fmt.Printf("Saldo akhir: Rp %d\n", saldo)
}