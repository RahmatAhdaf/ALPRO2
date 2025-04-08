package main

import (
	"fmt"
	"math"
)

func hitungRataRata(data []int) float64 {
	sum := 0
	for _, v := range data {
		sum += v
	}
	return float64(sum) / float64(len(data))
}

func hitungSimpanganBaku(data []int) float64 {
	rata := hitungRataRata(data)
	var total float64
	for _, v := range data {
		total += math.Pow(float64(v)-rata, 2)
	}
	return math.Sqrt(total / float64(len(data)))
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	fmt.Printf("\nIsi array yang dimasukkan: %v\n", data)
	fmt.Printf("Rata-rata: %.2f\n", hitungRataRata(data))
	fmt.Printf("Simpangan baku: %.2f\n", hitungSimpanganBaku(data))

	// Indeks genap dan ganjil
	fmt.Print("\nElemen dengan indeks genap: ")
	for i := 0; i < len(data); i += 2 {
		fmt.Printf("%d ", data[i])
	}
	fmt.Print("\nElemen dengan indeks ganjil: ")
	for i := 1; i < len(data); i += 2 {
		fmt.Printf("%d ", data[i])
	}
	fmt.Println()

	// Kelipatan indeks x
	var x int
	fmt.Print("\nMasukkan nilai kelipatan x: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen dengan indeks kelipatan %d : ", x)
	for i := 0; i < len(data); i++ {
		if i%x == 0 {
			fmt.Printf("%d ", data[i])
		}
	}
	fmt.Println()

	// Hapus elemen pada indeks tertentu
	var del int
	fmt.Print("\nMasukkan indeks yang ingin dihapus: ")
	fmt.Scan(&del)

	if del >= 0 && del < len(data) {
		data = append(data[:del], data[del+1:]...)
		fmt.Printf("Array setelah penghapusan: %v\n", data)
		fmt.Printf("\nRata-rata setelah penghapusan: %.2f\n", hitungRataRata(data))
		fmt.Printf("Simpangan setelah penghapusan: %.2f\n", hitungSimpanganBaku(data))
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	// Frekuensi bilangan
	var target int
	fmt.Print("\nMasukkan bilangan untuk mencari frekuensi: ")
	fmt.Scan(&target)
	count := 0
	for _, v := range data {
		if v == target {
			count++
		}
	}
	fmt.Printf("Frekuensi bilangan %d dalam array: %d\n", target, count)
}
