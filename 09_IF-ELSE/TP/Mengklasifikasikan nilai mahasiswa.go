package main

import "fmt"

func main() {
    var nilaimahasiswa int

    fmt.Print("Masukkan nilai mahasiswa: ")
    fmt.Scan(&nilaimahasiswa)

    if nilaimahasiswa > 90 {
        fmt.Println("Indeks yang diperoleh: A")
    } else if nilaimahasiswa >= 80 {
        fmt.Println("Indeks yang diperoleh: AB")
    } else if nilaimahasiswa >= 70 {
        fmt.Println("Indeks yang diperoleh: B")
    } else {
        fmt.Println("Indeks yang diperoleh: C")
    }
}
