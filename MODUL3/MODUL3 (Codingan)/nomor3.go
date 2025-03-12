package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik
func hitungJarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
}

// Fungsi untuk mengecek apakah titik berada dalam lingkaran
func dalamLingkaran(pusatX, pusatY, jariJari, titikX, titikY int) bool {
	return hitungJarak(pusatX, pusatY, titikX, titikY) <= float64(jariJari)
}

func main() {
	var pusatX1, pusatY1, jariJari1 int
	var pusatX2, pusatY2, jariJari2 int
	var titikX, titikY int

	// Input untuk lingkaran 1
	fmt.Print("Masukkan pusat dan jari-jari lingkaran 1 (pusatX1 pusatY1 jariJari1): ")
	fmt.Scan(&pusatX1, &pusatY1, &jariJari1)

	// Input untuk lingkaran 2
	fmt.Print("Masukkan pusat dan jari-jari lingkaran 2 (pusatX2 pusatY2 jariJari2): ")
	fmt.Scan(&pusatX2, &pusatY2, &jariJari2)

	// Input untuk titik
	fmt.Print("Masukkan koordinat titik (titikX titikY): ")
	fmt.Scan(&titikX, &titikY)

	// Cek posisi titik terhadap lingkaran
	adaDiLingkaran1 := dalamLingkaran(pusatX1, pusatY1, jariJari1, titikX, titikY)
	adaDiLingkaran2 := dalamLingkaran(pusatX2, pusatY2, jariJari2, titikX, titikY)

	// Menentukan keluaran berdasarkan kondisi
	if adaDiLingkaran1 && adaDiLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if adaDiLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if adaDiLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
