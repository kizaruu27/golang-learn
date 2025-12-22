package main

import "fmt"

func main() {
	sum, times := calculateStats(10, 2)
	fmt.Println("Hasil penjumlahan: ", sum)
	fmt.Println("Hasil perkalian: ", times)
}

func calculateStats(a, b int) (int, int) {
	sumValue := a + b
	timesValue := a * b

	return sumValue, timesValue
}