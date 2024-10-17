package main

import(
	"fmt"
)

func main(){
	var a,b int

	fmt.Println("Masukkan bairis bilangan kesatu : ")
	fmt.Scan(&a)

	fmt.Println("Masukkan bairis bilangan kedua : ")
	fmt.Scan(&b)

	for i:=a; i<=b; i++{
		fmt.Println(i)
	}
}
