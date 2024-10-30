package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&n)

	if n < 0 && n%2 == 0 {
		fmt.Println("genap negatif")
	} 
		fmt.Println("bukan")
}
