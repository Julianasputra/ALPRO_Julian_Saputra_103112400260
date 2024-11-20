package main

import "fmt"

func main() {

	var b, prima int
	var hasil bool

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&b)

	if b > 1 {
		fmt.Print("Faktor: ")

		for i := 1; i <= b; i++ {
			if b%i == 0 {
				fmt.Print(i, " ")
				prima++
			}
		}

		hasil = prima == 2
		fmt.Println()
		fmt.Println("Prima:", hasil)
	} else {
		fmt.Println("Masukkan bilangan bulat lebih dari 1!")
	}
}
