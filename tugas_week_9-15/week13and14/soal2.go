package main

import "fmt"

func decimalToBinary(decimal int) string {
        var binary string

        
        for decimal > 0 {
                remainder := decimal % 2
                binary = fmt.Sprintf("%d%s", remainder, binary)
                decimal /= 2
        }

        if binary == "" {
                return "0"
        }
        return binary
}

func main() {
        var decimal int
        fmt.Print("Masukkan bilangan desimal: ")
        fmt.Scan(&decimal)

        binary := decimalToBinary(decimal)
        fmt.Println("Bilangan biner: ", binary)
}