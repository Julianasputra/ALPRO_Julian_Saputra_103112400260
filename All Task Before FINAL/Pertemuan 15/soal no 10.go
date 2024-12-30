package main

import "fmt"

func main() {

	var j, y, bilangan, total int

	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	total = 0

	for j = 1; j <= 9; j++ {
		fmt.Printf("Masukkan bilangan ke-%d: ", j)
		fmt.Scan(&bilangan)
		total += bilangan
	}

	if total >= y*5 {
		fmt.Printf("Median bernilai %d\n", y)
	} else {
		fmt.Println("Median bernilai 0")
	}
}
