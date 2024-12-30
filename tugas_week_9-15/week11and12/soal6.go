package main

import "fmt"

func main() {
    var kapasitasTangki, volumeEmber int
    var isPenuh bool

    fmt.Scan(&kapasitasTangki)

    for i := 0; i < kapasitasTangki; i++ {
        fmt.Scan(&volumeEmber)
        kapasitasTangki -= volumeEmber

        if kapasitasTangki <= 0 {
            isPenuh = true
            break
        } else {
            isPenuh = false
        }

        fmt.Println(isPenuh)
    }
}