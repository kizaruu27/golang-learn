package main

import "fmt"

func main() {
	myNumbers := [6]int{10, 8, 5, 2, 6, 4}
	totalSum := 0

	for i, num := range myNumbers {
		fmt.Printf("Index %d: %d \n", i, num)
		totalSum += num
	}
	fmt.Println("Total:", totalSum)
}
