package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Print("Masukkan sisi pertama: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan sisi kedua: ")
	fmt.Scan(&b)

	fmt.Print("Masukkan sisi ketiga: ")
	fmt.Scan(&c)

	if a+b > c && a+c > b && b+c > a {
		if a == b && b == c {
			fmt.Println("Segitiga Sama Sisi")
		} else if a == b || b == c || a == c {
			fmt.Println("Segitiga Sama Kaki")
		} else {
			fmt.Println("Segitiga Sembarang")
		}
	} else {
		fmt.Println("Bukan Segitiga")
	}
}
