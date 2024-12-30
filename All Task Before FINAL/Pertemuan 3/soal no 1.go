package main

import "fmt"

func main() {
    var n int

    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scan(&n)

    if n%2 != 0 {
        fmt.Printf("true - %d adalah bilangan ganjil\n", n)
    } else {
        fmt.Printf("false - %d adalah bilangan genap\n", n)
    }
}
