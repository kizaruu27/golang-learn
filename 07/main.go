package main

import "fmt"

func main() {
	score := 100

	switch {
	case score == 100:
		fmt.Println("Nilai Sempurna")
	case score >= 80:
		fmt.Println("Nilai Bagus")
	case score >= 60:
		fmt.Println("Nilai Cukup")
	default:
		fmt.Println("Nilai Kurang")
	}
}
