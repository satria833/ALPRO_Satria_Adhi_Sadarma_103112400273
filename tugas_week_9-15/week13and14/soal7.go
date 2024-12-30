package main

import (
        "fmt"
        "strconv"
        
)

func main() {
        var input string
        var temps []int
        var sum, count int

        fmt.Println("Masukkan suhu (akhiri dengan dua angka 0 berturut-turut):")

        for {
                fmt.Scan(&input)
                temp, err := strconv.Atoi(input)
                if err != nil || temp == 0 {
                        if len(temps) > 0 && temps[len(temps)-1] == 0 {
                                break
                        }
                        continue
                }
                temps = append(temps, temp)
                sum += temp
                count++
            }

        if len(temps) == 0 {
                fmt.Println("Tidak ada data suhu yang valid.")
                return
        }

        maxTemp := temps[0]
        minTemp := temps[0]
        for _, temp := range temps {
                if temp > maxTemp {
                        maxTemp = temp
                }
                if temp < minTemp {
                        minTemp = temp
                }
        }

        avgTemp := float64(sum) / float64(count)

        fmt.Printf("Tertinggi: %d\n", maxTemp)
        fmt.Printf("Terendah: %d\n", minTemp)
        fmt.Printf("Rata-rata: %.3f\n", avgTemp)
}