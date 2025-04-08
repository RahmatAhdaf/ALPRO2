package main

import "fmt"

const NMAX = 127
type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var huruf rune
	*n = 0
	fmt.Print("Teks : ")
	for {
		fmt.Scanf("%c", &huruf)
		if huruf == '.' || *n >= NMAX {
			break
		}
		if huruf != ' ' && huruf != '\n' {
			t[*n] = huruf
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikkanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	var temp tabel
	for i := 0; i < n; i++ {
		temp[i] = t[i]
	}
	var reversed tabel = temp
	balikkanArray(&reversed, n)

	for i := 0; i < n; i++ {
		if temp[i] != reversed[i] {
			return false
		}
	}
	return true
}

func main() {
	var t tabel
	var n int

	isiArray(&t, &n)

	fmt.Print("Reverse teks : ")
	var tCopy tabel
	for i := 0; i < n; i++ {
		tCopy[i] = t[i]
	}
	balikkanArray(&tCopy, n)
	cetakArray(tCopy, n)

	if palindrom(t, n) {
		fmt.Println("Palindrome : true")
	} else {
		fmt.Println("Palindrome : false")
	}
}
