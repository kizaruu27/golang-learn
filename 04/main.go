package main

import "fmt"

func main() {
	printSum(10, 10)
}

func printSum(a, b int) {
	sum := a + b
	fmt.Printf("Hasil penjumlahan: %d\n", sum)
}