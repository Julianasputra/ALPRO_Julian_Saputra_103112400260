package main

import (
	"fmt"
)

func main() {
	var JumlahBarang int
	fmt.Print("Masukkan jumlah barang yang dibeli : ")
	fmt.Scan(&JumlahBarang)

	totalPoin := 0

	if JumlahBarang > 5 {
		totalPoin = (5 * 10) + ((JumlahBarang - 5) * 15)
	} else {
		totalPoin = JumlahBarang * 10
	}

	fmt.Printf("Total poin yang didapatkan : %d poin\n", totalPoin)
}
