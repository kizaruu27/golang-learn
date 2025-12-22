package main

import "fmt"

func main() {
	price := 100000
	discount := 50
	tax := 12.0
	var pricePointer *int = &price
	fmt.Printf("Harga awal: %d\n", price)

	// Discount price
	discountCut := price * discount / 100
	fmt.Printf("Nilai diskon: %d\n", discountCut)
	nettPrice := price - int(discountCut)
	fmt.Printf("Harga setelah diskon: %d\n", nettPrice)
	
	// Tax price
	taxValue := float64(nettPrice) * tax / 100
	fmt.Printf("Nilai pajak: %.2f\n", taxValue)
	
	*pricePointer = nettPrice + int(taxValue)
	fmt.Println("Harga akhir: ", price)
}