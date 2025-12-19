package main

import "fmt"

func main() {
	num := 5
	num++
	num++
	fmt.Printf("Hasil setelah increament sebanyak 2 kali: %d\n", num)

	num--
	fmt.Printf("Hasil setelah decreament sebanyak 1 kali: %d\n", num )
}