package main

import "fmt"

func main() {
	var n, exponen int

	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&exponen)

	result := 1 
	for i := 0; i < exponen; i++ {
		result *= n
	}

	fmt.Printf("%d dipangkatkan %d adalah %d\n", n, exponen, result)
}
