package main

import "fmt"

func cetakBintang(n, i int) {
    if i > n {
        return
    }
    
    for j := 0; j < i; j++ {
        fmt.Print("*")
    }
    fmt.Println()
    
    cetakBintang(n, i+1)
}

func main() {
    var n int
    fmt.Print("Masukkan jumlah baris: ")
    fmt.Scan(&n)
    
    cetakBintang(n, 1)
}
