package main

import "fmt"

func main() {
    // Variabel
    var bilangan int
	var hasil int

    fmt.Print("Masukan sebuah bilangan: ")
    fmt.Scan(&bilangan)

    switch {
    case bilangan%2 != 0 && bilangan%5 != 0 && bilangan%10 != 0:
        hasil = bilangan + (bilangan + 1)
        fmt.Printf("Kategori: Bilangan Ganjil\n")
        fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", bilangan, bilangan+1, hasil)
    case bilangan%2 == 0 && bilangan%5 != 0 && bilangan%10 != 0:
        hasil = bilangan * (bilangan + 1)
        fmt.Printf("Kategori: Bilangan Genap\n")
        fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", bilangan, bilangan+1, hasil)
    case bilangan%5 == 0 && bilangan%10 != 0:
        hasil = bilangan * bilangan
        fmt.Printf("Kategori: Bilangan Kelipatan 5\n")
        fmt.Printf("Hasil kuadrat dari %d ^2 = %d\n", bilangan, hasil)
    case bilangan%10 == 0:
        hasil = bilangan / 10
        fmt.Printf("Kategori: Bilangan Kelipatan 10\n")
        fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", bilangan, hasil)
    default:
        fmt.Println("Tidak ada kategori yang sesuai")
    }
}