package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat  Titik
	radius int
}

// Fungsi untuk menghitung jarak antara dua titik
func jarak(p, q Titik) float64 {
	dx := float64(p.x - q.x)
	dy := float64(p.y - q.y)
	return math.Sqrt(dx*dx + dy*dy)
}

// Fungsi untuk mengecek apakah sebuah titik ada di dalam lingkaran
func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) < float64(c.radius)
}

func main() {
	var x1, y1, r1 int
	var x2, y2, r2 int
	var px, py int

	// Input sesuai format di gambar
	fmt.Scan(&x1, &y1, &r1)
	fmt.Scan(&x2, &y2, &r2)
	fmt.Scan(&px, &py)

	// Buat lingkaran dan titik
	ling1 := Lingkaran{Titik{x1, y1}, r1}
	ling2 := Lingkaran{Titik{x2, y2}, r2}
	titik := Titik{px, py}

	// Cek posisi titik
	in1 := didalam(ling1, titik)
	in2 := didalam(ling2, titik)

	// Cetak output sesuai ketentuan
	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
