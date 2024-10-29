package main

import "fmt"

func main() {
    var a int

    fmt.Print("Masukkan sebuah angka: ")
    fmt.Scan(&a)

    if a%2 == 0 {
        fmt.Println("Angka adalah Genap.")
    } else {
        fmt.Println("Angka adalah Ganjil.")
    }
}
