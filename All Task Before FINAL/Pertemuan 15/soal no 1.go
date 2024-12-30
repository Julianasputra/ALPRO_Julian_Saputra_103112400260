package main

import "fmt"

func main() {

	var jDatang, mDatang, durasi int
	var datang, jSelesai, mSelesai int

	fmt.Print("Masukkan jam kedatangan (jDatang): ")
	fmt.Scan(&jDatang)
	fmt.Print("Masukkan menit kedatangan (mDatang): ")
	fmt.Scan(&mDatang)
	fmt.Print("Masukkan durasi layanan (dalam menit): ")
	fmt.Scan(&durasi)

	datang = jDatang*60 + mDatang + durasi 
	jSelesai = datang / 60                 
	mSelesai = datang % 60                 

	if jSelesai < 20 || (jSelesai == 20 && mSelesai == 0) {
		fmt.Printf("Layanan selesai pada pukul %02d:%02d\n", jSelesai, mSelesai)
	} else {
		fmt.Println("Silahkan datang kembali besok")
	}
}
