package main

import (
	"fmt"
)

func main() {
	
	const falsPerQirat = 6
	const dirhamPerFals = 10
	const dinarPerDirham = 10

	var qirat int
	fmt.Print("Masukkan jumlah uang dalam qirat: ")
	fmt.Scan(&qirat)

	fals := qirat / falsPerQirat
	sisaQirat := qirat % falsPerQirat

	dirham := fals / dirhamPerFals
	sisaFals := fals % dirhamPerFals

	dinar := dirham / dinarPerDirham
	sisaDirham := dirham % dinarPerDirham

	fmt.Printf("Hasil konversi:\n")
	fmt.Printf("Dinar: %d\n", dinar)
	fmt.Printf("Dirham: %d\n", sisaDirham)
	fmt.Printf("Fals: %d\n", sisaFals)
	fmt.Printf("Qirat: %d\n", sisaQirat)
}
