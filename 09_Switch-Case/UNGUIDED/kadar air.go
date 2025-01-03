package main

import "fmt"

func main() {
    var kadar float64

    fmt.Print("Masukan kadar pH: ")
    fmt.Scan(&kadar)

    switch {
    case kadar < 0 || kadar > 14:
        fmt.Println("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14")
    case kadar >= 6.5 && kadar <= 8.6:
        fmt.Println("Air layak minum")
    default:
        fmt.Println("Air tidak layak minum")
    }
}