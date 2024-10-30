package main

import "fmt"

func main() {
    var n int
    var sum int = 0

    fmt.Print("Masukkan bilangan bulat positif n: ")
    fmt.Scan(&n)

    for i := 1; i <= n; i++ {
        sum += i 
    }

    fmt.Println("Hasil penjumlahan dari 1 sampai", n, "adalah:", sum)
}
