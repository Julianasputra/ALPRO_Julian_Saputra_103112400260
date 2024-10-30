package main

import "fmt"

func main() {
    var a int

    fmt.Print("Masukkan bilangan bulat : ")
    fmt.Scan(&a)

    if a > 0 {
        fmt.Println("Positif")
    } else{
        fmt.Println("Bukan Positif")
	}
}
