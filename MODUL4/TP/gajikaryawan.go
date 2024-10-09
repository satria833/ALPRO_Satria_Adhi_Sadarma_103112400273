package main

import "fmt"

func main() {
	var hoursPerWeek, hourlyRate, overtimeHours, normalHours float64
	var weeklySalary, overtimePay, monthlySalary float64

	// Meminta input dari pengguna
	fmt.Print("Masukkan jumlah jam kerja dalam seminggu: ")
	fmt.Scanln(&hoursPerWeek)

	fmt.Print("Masukkan upah per jam: ")
	fmt.Scanln(&hourlyRate)

	// Jika jam kerja lebih dari 40, hitung lembur
	if hoursPerWeek > 40 {
		normalHours = 40
		overtimeHours = hoursPerWeek - 40
		overtimePay = overtimeHours * hourlyRate * 1.5
	} else {
		normalHours = hoursPerWeek
		overtimePay = 0
	}

	// Menghitung gaji mingguan
	weeklySalary = normalHours * hourlyRate

	// Menghitung gaji bulanan (4 minggu)
	monthlySalary = (weeklySalary + overtimePay) * 4

	// Menampilkan total gaji bulanan
	fmt.Printf("Total gaji bulanan: Rp %.2f\n", monthlySalary)
}
