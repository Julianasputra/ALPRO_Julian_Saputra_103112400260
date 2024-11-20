package main

import "fmt"

func main() {
	var nama string

	fmt.Scan(&nama)

	switch nama {
	case "nepenthes", "drosera":
		fmt.Println("Termasuk tanaman karnivora")
		fmt.Println("Asli Indonesia")
	case "venus", "sarracenia":
		fmt.Println("Termasuk tanaman karnivora")
		fmt.Println("Tidak asli Indonesia")
	default:
		fmt.Println("Tidak termasuk tanaman karnivora")
	}
}
