package main

import "fmt"

func main() {
	city := "Bogor"

	switch city {
	case "Bandung", "Bogor", "Bekasi":
		fmt.Println("Jawa Barat")
	case "Jakarta":
		fmt.Println("DKI Jakarta")
	default:
		fmt.Println("Kota tidak dikenal")
	}
}
