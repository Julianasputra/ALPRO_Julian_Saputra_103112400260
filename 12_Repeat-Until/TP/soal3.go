package main

import "fmt"

func main() {
	var harga, total int

	for {
		fmt.Print("Masukan harga barang (ketik 0 untuk selesai): ")
		fmt.Scan(&harga)

		total += harga

		if harga == 0 {
			break
		}
	}
	fmt.Println("Total belanja Anda: ", total)
}
