package main

import "fmt"

func main() {
	var x, y int
	hasil := 0
	fmt.Print("Masukan nilai x dan y: ")
	fmt.Scan(&x, &y)

	if x < y {
		fmt.Println("Bilangan x harus lebih besar dari bilangan y.")
		return
	}
	for x >= y {
		x -= y
		hasil++
	}
	fmt.Println(hasil)
}
