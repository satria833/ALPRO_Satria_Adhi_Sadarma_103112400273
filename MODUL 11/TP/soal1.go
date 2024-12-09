package main

import (
	"fmt"
)

func main() {
	
	correctPassword := "password123"

	
	maxAttempts := 3

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		
		var enteredPassword string
		fmt.Print("Masukkan password: ")
		fmt.Scan(&enteredPassword)

		
		if enteredPassword == correctPassword {
			fmt.Println("Login berhasil!")
			return 
		} else {
			
			fmt.Printf("Password salah. Anda telah mencoba %d kali dari %d percobaan.\n", attempts, maxAttempts)
		}
	}

	
	fmt.Println("Login ditolak.")
}
