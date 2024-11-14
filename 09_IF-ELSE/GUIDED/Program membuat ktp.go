package main 

import "fmt"

func main(){
	var usia int
	var kk bool

	fmt.Print("Masukkan usia : ")
	fmt.Scan(&usia)

	fmt.Print("Apakah memiliki kk (true/false) : ")
	fmt.Scan(&kk)

	if usia >= 17 && kk{
		fmt.Println("bisa membuat ktp")
	}else{
		fmt.Println("tidak bisa membuat ktp")
	}
}