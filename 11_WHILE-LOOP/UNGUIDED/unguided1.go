package main

import "fmt"

func main() {
	var username, password string

	correctUsername := "Admin"
	correctPassword := "Admin1"

	jumlahPercobaan := 0

	for {
		fmt.Print("Masukan username dan password: ")
		fmt.Scan(&username, &password)

		if username == correctUsername && password == correctPassword {
			fmt.Printf("%d percobaan gagal login.\n", jumlahPercobaan)
			break
		} else {
			jumlahPercobaan++
		}
	}
}
