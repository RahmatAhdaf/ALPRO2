package main

import "fmt"

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Print("Masukkan nilai a, b, c, d: ")
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		p_ac := permutation(a, c)
		c_ac := combination(a, c)
		p_bd := permutation(b, d)
		c_bd := combination(b, d)

		fmt.Println(p_ac, c_ac)
		fmt.Println(p_bd, c_bd)
	} else {
		fmt.Println("Input tidak memenuhi syarat a >= c dan b >= d")
	}
}
