package main

import "fmt"

func main() {
	fruits := [5]string{"Apel", "Pisang", "Semangka", "Rambutan", "Alpukat"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index %d: %s \n", i, fruits[i])
	}
}
