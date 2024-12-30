package main

import "fmt"

func main() {
	var panjang, lebar, luas, keliling float64

	fmt.Print("Masukkan panjang: ")
	fmt.Scanln(&panjang)
	fmt.Print("Masukkan lebar: ")
	fmt.Scanln(&lebar)

	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)

	fmt.Println("\nInformasi Persegi Panjang:")
	fmt.Println("Panjang:", panjang)
	fmt.Println("Lebar:", lebar)
	fmt.Println("Luas:", luas)
	fmt.Println("Keliling:", keliling)
}
