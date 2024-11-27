package main

import "fmt"

func main() {
	var password string
	var passwordbenar = "1234"

	for i := 1; i <= 3; i++ {
		fmt.Print("Masukan password terlebih dahulu: ")
		fmt.Scan(&password)

		if password == passwordbenar {
			fmt.Println("Anda berhasil Login!!!")
			break
		} else {
			fmt.Println("Password yang anda masukan salah, coba lagi!!!")
		}
		if i == 3 {
			fmt.Println("Login ditolak")
		}
	}
}
