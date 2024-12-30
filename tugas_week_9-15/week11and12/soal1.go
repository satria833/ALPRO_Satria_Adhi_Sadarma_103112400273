package main

import "fmt"

func main() {
        var username, password string
        var gagalLogin int

        for {
                fmt.Print("Masukkan username: ")
                fmt.Scanln(&username)
                fmt.Print("Masukkan password: ")
                fmt.Scanln(&password)

                if username == "admin" && password == "admin" {
                        fmt.Println("Login berhasil!")
                        break
                } else {
                        fmt.Println("Username atau password salah. Silakan coba lagi.")
                        gagalLogin++
                }
        }

        fmt.Printf("Jumlah gagal login: %d\n", gagalLogin)
}