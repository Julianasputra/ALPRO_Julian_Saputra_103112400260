package main

import "fmt"

func main() {
	var angka int

	for {
		fmt.Print("Tebak angka (1-10): ")
		fmt.Scan(&angka)

		if angka == 7 {
			fmt.Println("Selamat, tebakan Anda benar!.")
			break
		} else {
			fmt.Println("Tebakan Anda salah, coba lagi.")
		}
	}
}
