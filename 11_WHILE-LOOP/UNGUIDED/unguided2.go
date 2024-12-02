package main

import "fmt"

func main() {
	var n, digit int

	fmt.Print("Masukan bilangan bulat positif: ")
	fmt.Scan(&n)

	for n > 0 {
		digit = n % 10
		fmt.Println(digit)
		n = n / 10
	}
}
