package main

import "fmt"

func main() {
	var bil int
	var prima bool

	fmt.Print("Masukan bilangan bulat: ")
	fmt.Scan(&bil)

	for i := 2; i <= bil; i++ {
		prima = true

		for a := 2; a*a <= i; a++ {

			if i%a == 0 {
				prima = false
				break
			}
		}
		if prima {
			fmt.Print(i, " ")
		}
	}
}
