package main

import "fmt"

func main() {
	var x int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&x)

	if x <= 0 {
	}

	total := 0
	fmt.Print("Digit: ")

	for x > 0 {
		digit := x % 10
		fmt.Print(digit, " ")
		total += digit
		x /= 10
	}

	fmt.Println()

	fmt.Println("Total:", total)
}
