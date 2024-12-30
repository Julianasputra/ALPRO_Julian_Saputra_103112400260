package main

import "fmt"

func main() {
	var n int
	var s string

	fmt.Print("Masukkan jumlah pengulangan (n): ")
	fmt.Scanln(&n)

	fmt.Print("Masukkan string: ")
	fmt.Scanln(&s)

	for i := 0; i < n; i++ {
		fmt.Println(s)
	}
}
