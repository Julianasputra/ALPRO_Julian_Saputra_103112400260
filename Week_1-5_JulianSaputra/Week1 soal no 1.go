package main

import "fmt"

func main() {
	var panjang, lebar, tinggi float64

	fmt.Print("Masukkan panjang lemari (dalam meter): ")
	fmt.Scan(&panjang)

	fmt.Print("Masukkan lebar lemari (dalam meter): ")
	fmt.Scan(&lebar)

	fmt.Print("Masukkan tinggi lemari (dalam meter): ")
	fmt.Scan(&tinggi)

	volume := panjang * lebar * tinggi

	fmt.Println("Volume lemari adalah:", volume, "meter")
}
