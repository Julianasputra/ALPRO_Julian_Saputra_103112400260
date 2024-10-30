package main

import "fmt"

func main() {
	var jumlahOrang int

	fmt.Print("Masukkan jumlah orang : ")
	fmt.Scan(&jumlahOrang)

	jumlahMotor := jumlahOrang / 2 

	if jumlahOrang%2 != 0 {
		jumlahMotor++
	}

	fmt.Printf("Jumlah motor yang diperlukan: %d\n", jumlahMotor)
}
