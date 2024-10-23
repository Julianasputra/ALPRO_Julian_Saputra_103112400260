package main

import (
	"fmt"
)

func main() {
	var x, y, hasil int

	fmt.Printf("Masukkan bilangan bulat positif pertama : ")
	fmt.Scan(&x)

	fmt.Printf("Masukkan bilangan bulat positif kedua : ")
	fmt.Scan(&y)

	for i := x; i <= y; i++ {
		hasil = hasil + i
	}
	fmt.Println("Hasil : ", hasil)
}
