package main

import "fmt"

func main() {
	var jeniskendaraan string
	var durasiparkir, tarif int

	fmt.Print("Masukan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&jeniskendaraan)

	fmt.Print("Masukan durasi parkir (dalam jam): ")
	fmt.Scan(&durasiparkir)

	switch {
	case durasiparkir < 1:
		durasiparkir = 1
	}
	switch jeniskendaraan {
	case "motor":
		tarif = 2000 * durasiparkir
	case "mobil":
		tarif = 5000 * durasiparkir
	case "truk":
		tarif = 8000 * durasiparkir
	default:
		fmt.Println("Jenis kendaraan tidak valid")
	}
	fmt.Printf("Tarif parkir: Rp%d\n", tarif)
}
