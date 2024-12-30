package main

import "fmt"

func main() {
	var a, b, sum int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scanln(&b)

	if a%2 == 0 {
		a++
	}

	for i := a; i <= b; i += 2 {
		sum += i
	}

	fmt.Println("Hasil penjumlahan bilangan ganjil:", sum)
}
