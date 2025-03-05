package main

import "fmt"

func main() {
	var jejari float64
	fmt.Print("Jejari = ")
	fmt.Scan(&jejari)

	volume := (4.0 / 3.0) * 3.1415926535 * (jejari * jejari * jejari)
	luas := 4 * 3.1415926535 * (jejari * jejari)

	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", jejari, volume, luas)
}
