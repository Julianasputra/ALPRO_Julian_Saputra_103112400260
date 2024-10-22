package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)

	if n > 0 {
		for i := 1; i <= n; i++ {
			fmt.Printf("Kuadrat dari %d adalah %d\n", i, i*i)
		}
	} else {
		fmt.Println("Harap masukkan bilangan bulat positif!")
	}
}
