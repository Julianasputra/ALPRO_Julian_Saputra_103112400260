package main

import "fmt"

func main() {
	var n int
	var kata string

	fmt.Print("Masukkan jumlah pencetakan: ")
	fmt.Scanln(&n)

	fmt.Print("Masukkan kata yang akan dicetak: ")
	fmt.Scanln(&kata)

	for i := 0; i < n; i++ {
		fmt.Println("Hasil : ",kata)
	}
}
