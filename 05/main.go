package main

import "fmt"

func main() {
	totalPrice := calculateDiscount(100000, 50)
	fmt.Println("Total harga setelah diskon: ", totalPrice)
}

func calculateDiscount(price, discount int) int {
	discountValue := price * discount / 100
	totalPrice := price - discountValue

	return totalPrice
}