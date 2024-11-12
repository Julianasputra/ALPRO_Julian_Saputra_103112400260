package main

import "fmt"

func main() {
	var masukkan string

	fmt.Print("Masukkan sebuah huruf besar atau kecil : ")
	fmt.Scan(&masukkan)

	if masukkan == "A" || masukkan == "I" || masukkan == "U" || masukkan == "E" || masukkan == "O" ||
	masukkan == "a" || masukkan == "i" || masukkan == "u" || masukkan == "e" || masukkan == "o" {
		fmt.Println("Huruf Vokal")
	} else {
		fmt.Println("Huruf Konsonan")
	}
}
