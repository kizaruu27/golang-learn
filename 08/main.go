package main

import "fmt"

func main() {
	sumTotal, _, _ := calculateAll(10, 10)
	fmt.Println("Hasil penjumlahan:", sumTotal)
}

func calculateAll(a, b int) (int, int, int) {
	penjumlahan := a + b
	pengurangan := a - b
	perkalian := a * b

	return penjumlahan, pengurangan, perkalian
}