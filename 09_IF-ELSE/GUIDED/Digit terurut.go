package main

import "fmt"

func main() {
	var angka int

	fmt.Print("Bilangan 4 digit : ")
	fmt.Scan(&angka)

	if angka < 1000 || angka > 9999 {
		fmt.Println("Bilangan harus 4 digit antara 1000 sampai 9999")
	}
	digit1 := angka / 1000
	digit2 := (angka / 100) % 10
	digit3 := (angka / 10) % 10
	digit4 := angka % 10

	if digit1 < digit2 && digit2 < digit3 && digit3 < digit4 {
		fmt.Println("Digit terurut membesar")
	} else if digit1 > digit2 && digit2 > digit3 && digit3 > digit4 {
		fmt.Println("Digit terurut mengecil")
	} else {
		fmt.Println("Digit tidak terurut")
	}
}
