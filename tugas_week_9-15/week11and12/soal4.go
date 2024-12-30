package main

import "fmt"

func main() {
        var n, m, x, y, cangkir int

        fmt.Scan(&n, &m, &x, &y)

        
        cangkirGula := n / x

        
        cangkirKopi := m / y

        
        cangkir = min(cangkirGula, cangkirKopi)

        fmt.Println(cangkir)
}

func min(a, b int) int {
        if a < b {
                return a
        }
        return b
}