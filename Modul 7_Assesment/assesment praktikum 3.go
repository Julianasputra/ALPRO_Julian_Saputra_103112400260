package main

import (
	"fmt"
)

func main() {
	var dinar, dirham, fals, qirat, sisa int

	fmt.Print("Masukkan nilai uang dalam satuan qirat : ")
	fmt.Scan(&qirat)

	dinar = qirat / (10 * 10 * 6)
	sisa = qirat % (10 * 10 * 6)

	dirham = sisa / (10 * 6)
	sisa = sisa % (10 * 6)

	fals = sisa / 6
	sisa = sisa % 6

	qirat = sisa

	fmt.Println("Hasil : ", dinar, dirham, fals, qirat)
}
