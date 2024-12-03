package main

import "fmt"

func main() {
	var a string

	for {
		fmt.Print("Masukan kata: ")
		fmt.Scan(&a)

		if a == "telkom" {
			fmt.Println("Program selesai")
			break
		}
		fmt.Println("Anda mengetik: ", a)
	}
}
