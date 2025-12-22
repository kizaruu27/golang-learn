package main

import "fmt"

func main() {
	area := calculateArea(20, 2)
	fmt.Println("Luas area:", area)
}

func calculateArea(width, length int) (area int) {
	area = width * length
	return
}