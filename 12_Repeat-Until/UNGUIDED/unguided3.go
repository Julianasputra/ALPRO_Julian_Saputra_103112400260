package main

import "fmt"

func main() {
	var target, donation, totalDonations, donorCount int

	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	for totalDonations < target {
		fmt.Printf("Masukkan jumlah donasi donatur %d: ", donorCount+1)
		fmt.Scan(&donation)

		totalDonations += donation
		donorCount++

		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", donorCount, donation, totalDonations)
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", totalDonations, donorCount)
}
