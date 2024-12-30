package main

import "fmt"

func main() {
    var h1, m1, h2, m2, ih, mm int

    fmt.Print("Jam mulai parkir (1-12): ")
    fmt.Scan(&h1)
    fmt.Print("Menit mulai parkir (0-59): ")
    fmt.Scan(&m1)
    fmt.Print("Jam selesai parkir (1-12): ")
    fmt.Scan(&h2)
    fmt.Print("Menit selesai parkir (0-59): ")
    fmt.Scan(&m2)

    if h2 > 5 || (h2 == 5 && m2 > 0) {
        fmt.Println("Kendaraan keluar setelah tempat wisata tutup.")
        return
    }

    totalMenit := h2*60 + m2 - h1*60 - m1

    if totalMenit < 0 {
        totalMenit += 12 * 60
    }

    ih = totalMenit / 60
    mm = totalMenit % 60

    fmt.Printf("Durasi parkir: %d jam %d menit\n", ih, mm)
}