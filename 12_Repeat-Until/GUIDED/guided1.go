package main

import "fmt"

func main(){
	var a string
	var kata int

	fmt.Print("Masukkan kata : ")
	fmt.Scan(&a,&kata)

	counter := 0
	for selesai := false; !selesai; {
		fmt.Println(a)
		counter ++
		selesai = (counter >= kata)
	}
}