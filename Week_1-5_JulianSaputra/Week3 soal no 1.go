package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	if n%2 != 0 {
		fmt.Println("true") 
	} else {
		fmt.Println("false")
	}
}
