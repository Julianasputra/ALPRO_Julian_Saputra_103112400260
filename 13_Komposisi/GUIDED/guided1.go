package main

import "fmt"

func main(){
	var bil int

	fmt.Print("Masukkan bIlangan : ")
	fmt.Scan(&bil)

	for i:=1; i <= bil; i += 2{
		fmt.Println(i)
	}
}