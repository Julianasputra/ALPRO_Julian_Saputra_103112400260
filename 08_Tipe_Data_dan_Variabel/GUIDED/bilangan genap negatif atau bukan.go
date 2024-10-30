package main

import "fmt"

func main() {
    var i int
	var j bool

    fmt.Print("Masukkan bilangan bulat : ")
    fmt.Scan(&i)

    if i < 0 && i%2 == 0 {
        j = true
    } 
        fmt.Println(j)
	}

