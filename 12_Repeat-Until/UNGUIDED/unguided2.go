package main

import "fmt"

func main() {
	var input float64

	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&input)

	akhir := int(input)
	if float64(akhir) < input {
		akhir++
	}

	fmt.Println("Hasil penjumlahan tiap perulangan:")
	current := input
	for current < float64(akhir) {
		current += 0.1
		fmt.Printf("%.1f\n", current)
	}
}
