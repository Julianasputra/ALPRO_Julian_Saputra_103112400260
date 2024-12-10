package main

import "fmt"

func main() {
	var bil, jumlah int

	fmt.Print("Masukan bilangan: ")
	fmt.Scan(&bil)

	for i := 1; i < bil; i++ {
		if bil%i == 0 {
			jumlah += i
		}
	}
	if jumlah == bil {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}
