package main

import "fmt"

func main() {
	addResult := calculate(10, 2, add)
	multiplyResult := calculate(10, 2, multiply)
	fmt.Println("Hasil add:", addResult)
	fmt.Println("Hasil multiply:", multiplyResult)
}

func calculate(a, b int, operation func(int, int) int) int {
	result := operation(a, b)
	return result
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}