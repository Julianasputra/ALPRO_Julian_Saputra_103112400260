package main

import "fmt"

func main() {
	var p, l int

	fmt.Print("Masukkan panjang (p): ")
	fmt.Scan(&p)
	fmt.Print("Masukkan lebar (l): ")
	fmt.Scan(&l)

	keliling := 2 * (p + l)
	luas := p * l

	fmt.Printf("Keliling: %d\n", keliling)
	fmt.Printf("Luas: %d\n", luas)
}
