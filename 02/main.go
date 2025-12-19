package main

import "fmt"

func main() {
	value1 := 10
	value2 := 77

	fmt.Printf("Hasil penjumlahan: %v\n", value1 + value2)
	fmt.Printf("Hasil pengurangan: %v\n", value2 - value1)
	// Parsing value dari float ke int
	fmt.Printf("Hasil perkalian: %v\n", int(value2) * int(value1))
	// Parsing dari int ke float
	fmt.Printf("Hasil pembagian: %v\n", float64(value2) / float64(value1))
}