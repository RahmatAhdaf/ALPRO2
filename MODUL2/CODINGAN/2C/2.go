// Code yang salah:
//package main
// 
//import "fmt"
// 
//func main() {
//	var nam float64
// 	var nmk string
// 	fmt.Print(“Nilai akhir mata kuliah: “)
//	fmt.Scanln(&nam)
// 	if nam > 80 {
// 		nam = “A”
// 	}
// 	if nam > 72.5 {
// 		nam = “AB”
//	}
// 	if nam > 65 {
//		nam = “B”
//	}
//	 if nam > 57.5 {
//		nam = “BC”
//	}
// 	if nam > 50 {
// 		nam = “C”
// 	}
//	 if nam > 40 {
// 		nam = “D”
//	} else if nam <= 40 {
//		nam = “E”
//	}
//	fmt.Println(“Nilai mata kuliah: “, nmk)
//}


// Jawab pertanyaan:
// Jika nilai NAM diberikan sebesar 80.1, program akan mengeluarkan "A", padahal seharusnya "AB", jadi program tidak sesuai dengan spesifikasi soal.
// Kesalahan utama adalah penggunaan if terpisah tanpa else if, sehingga setiap kondisi tetap dicek meskipun yang sebelumnya sudah terpenuhi. Selain itu, batas kategori tidak dibuat dengan benar, seperti pada nilai di atas 80 yang langsung masuk "A" tanpa memperhitungkan batas atasnya.
// Agar program benar, harus menggunakan else if supaya nilai hanya masuk ke satu kategori yang sesuai. Jika diperbaiki, masukan 93.5, 70.6, dan 49.5 akan menghasilkan keluaran "A", "B", dan "D" seperti yang diharapkan.



// Code yang sudah perbaiki:
package main

import "fmt"

func main() {
	var nam float64
	var nmk string

	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)

	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "AB"
	} else if nam > 65 {
		nmk = "B"
	} else if nam > 57.5 {
		nmk = "BC"
	} else if nam > 50 {
		nmk = "C"
	} else if nam > 40 {
		nmk = "D"
	} else {
		nmk = "E"
	}

	fmt.Println("Nilai mata kuliah:", nmk)
}


