package main

import "fmt"

func hitungSkor(waktu [8]int) (int, int) {
	totalWaktu, soalSelesai := 0, 0
	for _, w := range waktu {
		if w <= 300 {
			totalWaktu += w
			soalSelesai++
		}
	}
	return soalSelesai, totalWaktu
}

func main() {
	var nama, pemenang string
	var waktu [8]int
	maksSoal, minWaktu := -1, 999999

	for {
		fmt.Print("Masukkan nama peserta dan waktu pengerjaan soal (atau 'Selesai' untuk berhenti): ")
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		soalSelesai, totalWaktu := hitungSkor(waktu)

		if soalSelesai > maksSoal || (soalSelesai == maksSoal && totalWaktu < minWaktu) {
			pemenang, maksSoal, minWaktu = nama, soalSelesai, totalWaktu
		}
	}

	fmt.Printf("%s %d %d\n", pemenang, maksSoal, minWaktu)
}