package main

import "fmt"

func main() {
	totalSumary := sumAll(4, 4, 4, 4)
	fmt.Println("Hasil penjumlahan total:", totalSumary)
}

func sumAll(values ...int) int {
	var total int
	for i := 0; i < len(values); i++ {
		total += values[i]
	}

	return total
}