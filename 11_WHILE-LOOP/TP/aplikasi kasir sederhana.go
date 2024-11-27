package main

import "fmt"

func main() {
	var hargabarang, jumlahbarang int
	var totalbelanja float64
	var tambahbarang, namabarang string

	fmt.Println("<==== Aplikasi Kasir Sederhana ====>")
	for {
		fmt.Print("Masukan nama barang: ")
		fmt.Scan(&namabarang)

		fmt.Print("Masukan jumlah barang: ")
		fmt.Scan(&jumlahbarang)

		fmt.Print("Masukan harga barang: Rp.")
		fmt.Scan(&hargabarang)

		totalbelanja += float64(hargabarang) * float64(jumlahbarang)

		fmt.Print("Apakah ada tambahan barang lagi? (ya/tidak): ")
		fmt.Scan(&tambahbarang)

		if tambahbarang == "tidak" {
			break
		}
	}
	fmt.Printf("\nTotal belanja anda: Rp.%.2f\n", totalbelanja)
	fmt.Println("<=== Terima kasih sudah berbelanja!!! ===>")
}
