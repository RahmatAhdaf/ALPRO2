package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e int
	var x, y, z byte

	// Membaca 5 angka integer
	fmt.Scan(&a, &b, &c, &d, &e)
	fmt.Scanln()
	// Membaca 3 karakter tanpa spasi
	fmt.Scanf("%c%c%c", &x, &y, &z)

	// Mencetak karakter dari angka ASCII
	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)
	// Mencetak karakter setelahnya dalam ASCII
	fmt.Printf("%c%c%c\n", x+1, y+1, z+1)
}
