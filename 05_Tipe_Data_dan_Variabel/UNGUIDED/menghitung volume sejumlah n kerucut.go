package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Print("Masukkan jumlah kerucut (n): ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var r float64 
		var h float64 

		fmt.Print("Masukkan jari-jari alas kerucut ke-", i, ": ")
		fmt.Scan(&r)
		fmt.Print("Masukkan tinggi kerucut ke-", i, ": ")
		fmt.Scan(&h)

		volume := (1.0 / 3.0) * math.Pi * r * r * h

		fmt.Printf("Volume kerucut ke-%d adalah: %.2f\n", i, volume)
	}
}
