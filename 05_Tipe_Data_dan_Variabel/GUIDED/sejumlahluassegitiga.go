package main

import(
	"fmt"
)

func main(){
	var n,alas,tinggi int

	fmt.Println("Masukkan n : ")
	fmt.Scan(&n)


	for i:=1; i<=n; i++{
		
		fmt.Println("Masukkan alas : ")
		fmt.Scan(&alas)

		fmt.Println("Masukkan tinggi : ")
		fmt.Scan(&tinggi)

		luas:= alas*tinggi/2

		fmt.Println("Luas:",luas)	
	}
}