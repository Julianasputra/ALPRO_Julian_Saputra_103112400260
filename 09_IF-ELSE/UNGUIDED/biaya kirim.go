package main

import "fmt"

func main() {
	var berat, digit1, digit2, biayaKg, sisaBiaya, totalBiaya int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	digit1 = berat / 1000
	digit2 = berat % 1000

	biayaKg = digit1 * 10000

	if digit2 >= 500 {
		sisaBiaya = digit2 * 5
	} else {
		sisaBiaya = digit2 * 15 
	}

	if digit1 > 10 {
		sisaBiaya = 0
	}

	totalBiaya = biayaKg + sisaBiaya

	fmt.Printf("Detail berat: %d kg + %d gr\n", digit1, digit2)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, sisaBiaya)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
