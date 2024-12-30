package main

import "fmt"

func main() {
	var input string

	fmt.Print("Masukkan satu karakter: ")
	fmt.Scan(&input)

	switch input {
	case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
		fmt.Println("bukan konsonan")
	default:
		if (input >= "a" && input <= "z") || (input >= "A" && input <= "Z") {
			fmt.Println("konsonan")
		} else {
			fmt.Println("bukan konsonan")
		}
	}
}
