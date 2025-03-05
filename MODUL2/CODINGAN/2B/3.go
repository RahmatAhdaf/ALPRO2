package main

import "fmt"

func main() {
	var berat1, berat2, total float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		// Berhenti jika salah satu kantong negatif atau total lebih dari 150 kg
		if berat1 < 0 || berat2 < 0 || total+berat1+berat2 > 150 {
			break
		}

		total += berat1 + berat2

		// Cek apakah selisih lebih dari 9 kg
		if berat1-berat2 >= -9 && berat1-berat2 <= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		}
	}

	fmt.Println("Proses selesai.")
}
