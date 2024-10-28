package main

import "fmt"

func main() {
	var p, l int

	fmt.Print("Masukkan panjang : ")
	fmt.Scan(&p)

	fmt.Print("Masukkan lebar : ")
	fmt.Scan(&l)

	luas := p * l

	keliling := 2 * (p + l)

	fmt.Println("Luas persegi panjang adalah:", luas)
	fmt.Println("Keliling persegi panjang adalah:", keliling)
}
