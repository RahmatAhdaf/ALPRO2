package main

import "fmt"


func turunNaik(n, i int) {
    if i == 1 {
        fmt.Print(i, " ")
        return
    }
    fmt.Print(i, " ")
    turunNaik(n, i-1)
    fmt.Print(i, " ")
}

func main() {
    var n int
    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&n)
    
    turunNaik(n, n)
    fmt.Println()
}