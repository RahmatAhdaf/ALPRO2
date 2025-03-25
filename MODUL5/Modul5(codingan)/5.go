package main

import "fmt"

func cetakGanjil(i, n int) {
    if i > n {
        return
    }
    fmt.Print(i, " ")
    cetakGanjil(i+2, n)
}

func main() {
    var n int
    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&n)
    
    cetakGanjil(1, n)
    fmt.Println()
}
