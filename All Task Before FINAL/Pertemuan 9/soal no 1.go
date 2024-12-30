package main

import "fmt"

func main() {
	var N, mobil int

	fmt.Scan(&N)

	if N <= 7 {
		mobil = 1
	} else if N <= 14 {
		mobil = 2
	} else if N <= 21 {
		mobil = 3
	} else {
		mobil = (N + 6) / 7
	}

	fmt.Println(mobil)
}
