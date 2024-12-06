package main

import "fmt"

func main(){
	var angka int
	var continueLoop bool

	for continueLoop = true; continueLoop; {
		fmt.Print("Masukkan angka : ")
		fmt.Scan(&angka)
		continueLoop = angka <= 0
	}
	fmt.Printf("%d adalah bilangan bulat positif\n", angka)
}