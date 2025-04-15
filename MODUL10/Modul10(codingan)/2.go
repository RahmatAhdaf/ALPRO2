package main

import "fmt"

func main() {
	var x, y int
	var berat [1000]float64

	fmt.Print("Masukkan jumlah ikan dan kapasitas wadah: ")
	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Printf("Masukkan berat ikan ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	jumlahWadah := (x + y - 1) / y
	var totalWadah [1000]float64

	idx := 0
	for i := 0; i < jumlahWadah; i++ {
		var total float64 = 0
		for j := 0; j < y && idx < x; j++ {
			total += berat[idx]
			idx++
		}
		totalWadah[i] = total
	}

	var totalSemua float64 = 0
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Total berat wadah %d: %.2f\n", i+1, totalWadah[i])
		totalSemua += totalWadah[i]
	}

	rata := totalSemua / float64(jumlahWadah)
	fmt.Printf("Rata-rata berat per wadah: %.2f\n", rata)
}
