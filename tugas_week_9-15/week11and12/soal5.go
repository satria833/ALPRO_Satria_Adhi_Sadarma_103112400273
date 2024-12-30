package main

import "fmt"

func isKonsekutif(x int) bool {
        prevDigit := -1
        for x > 0 {
                digit := x % 10
                if prevDigit != -1 && abs(prevDigit-digit) != 1 {
                        return false
                }
                prevDigit = digit
                x /= 10
            }
        return true
}

func abs(x int) int {
        if x < 0 {
                return -x
        }
        return x
    }

func main() {
        var x int
        fmt.Print("Masukkan bilangan: ")
        fmt.Scan(&x)

        if isKonsekutif(x) {
                fmt.Println("true")
        } else {
                fmt.Println("false")
        }
}