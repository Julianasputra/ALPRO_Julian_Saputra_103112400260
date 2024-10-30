package main

import "fmt"

func main(){
	var n int

	fmt.Print("Masukkan bilangan bulat : ")
	fmt.Scan(&n)

	if n < 0 {
		n = -n
	}
	fmt.Println("Hasil : ",n)
}
