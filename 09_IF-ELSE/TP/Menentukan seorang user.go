package main

import "fmt"

func main() {
	var umur int
	var kewarganegaraan string

	fmt.Print("Masukkan umur : ")
	fmt.Scan(&umur)

	fmt.Print("Masukkan kewarganegaraan (WNI atau WNA): ")
	fmt.Scan(&kewarganegaraan)

	if umur >= 17 {
		if kewarganegaraan == "WNI" {
			fmt.Println("Anda bisa mengikuti pemilu")
		} else {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena Anda bukan WNI")
		}
	} else {
		fmt.Println("Anda tidak bisa mengikuti pemilu karena umur Anda kurang dari 17 tahun")
	}
}
