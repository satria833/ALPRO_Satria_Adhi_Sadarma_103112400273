package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	var guess int

	fmt.Println("Selamat datang di permainan tebak angka!")
	fmt.Println("Saya sudah memilih sebuah angka antara 1 hingga 100.")
	fmt.Println("Anda memiliki 5 kesempatan untuk menebak.")

	for attempts := 1; attempts <= 5; attempts++ {
		fmt.Printf("Kesempatan ke-%d: Masukkan tebakan Anda: ", attempts)
		fmt.Scan(&guess)

		if guess < target {
			fmt.Println("Tebakan Anda terlalu kecil.")
		} else if guess > target {
			fmt.Println("Tebakan Anda terlalu besar.")
		} else {
			fmt.Printf("Selamat! Anda telah menebak angka %d dengan benar!\n", target)
			return
		}
	}

	fmt.Printf("Sayang sekali! Anda telah kehabisan kesempatan. Angka yang benar adalah %d.\n", target)
}