package main

import (
	"fmt"
)

func main() {
	var angka int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&angka)

	if angka < 0 {
	}

	digitCount := 0
	digit := angka

	for digit > 0 {
		digit /= 10
		digitCount++
	}

	fmt.Printf("Jumlah digit dari bilangan %d adalah %d.\n", angka, digitCount)
}
