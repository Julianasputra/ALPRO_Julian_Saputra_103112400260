package main

import "fmt"

func main() {
    var n int

    fmt.Print("Masukkan nilai ujian: ")
    fmt.Scan(&n)

    if n >= 70 {
        fmt.Println("Lulus")
    } else {
        fmt.Println("Tidak Lulus")
    }
}
