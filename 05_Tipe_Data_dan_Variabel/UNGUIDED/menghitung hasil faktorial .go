package main

import "fmt"

func main() {
	var b int

	fmt.Print("Masukkan bilangan bulat non negatif : ")
	fmt.Scan(&b)

	faktorial := 1 

	for i := 1; i <= b; i++ {
		faktorial *= i 
	}

	fmt.Printf("Faktorial dari %d adalah %d\n", b, faktorial)
}
