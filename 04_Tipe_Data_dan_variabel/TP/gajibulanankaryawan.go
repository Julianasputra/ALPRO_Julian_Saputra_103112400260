package main

import (
	"fmt"
)

func main() {
	var jamKerjaPerMinggu, upahPerJam float64

	fmt.Print("Masukkan jumlah jam kerja per minggu: ")
	fmt.Scan(&jamKerjaPerMinggu)
	fmt.Print("Masukkan upah per jam: ")
	fmt.Scan(&upahPerJam)

	var jamNormal, jamLembur float64
	if jamKerjaPerMinggu > 40 {
		jamNormal = 40
		jamLembur = jamKerjaPerMinggu - 40
	} else {
		jamNormal = jamKerjaPerMinggu
		jamLembur = 0
	}

	totalGaji := (jamNormal * upahPerJam * 4) + (jamLembur * 1.5 * upahPerJam * 4)

	fmt.Printf("Total gaji bulanan adalah: %.2f\n", totalGaji)
}
