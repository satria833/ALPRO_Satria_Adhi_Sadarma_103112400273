package main

import "fmt"

func main() {
	var umur int
	fmt.Print("Masukan Umur: ")
	fmt.Scan(&umur)

	var kewarganegaraan string
	fmt.Print("Masukan Kewarganegaraan WNI/WNA: ")
	fmt.Scan(&kewarganegaraan)

	if umur >= 17 && (kewarganegaraan == "WNI" || kewarganegaraan == "wni") {
		fmt.Println("Anda bisa mengikuti pemilu")
	} else {
		if umur < 17 {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena umur anda kurang dari 17 tahun.")
		}
		if kewarganegaraan != "WNI" && kewarganegaraan != "wni" {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena anda bukan WNI")
		}
	}
}