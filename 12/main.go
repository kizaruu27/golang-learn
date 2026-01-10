// --- Switch case tanpa nilai ---

package main

import "fmt"

func main() {
	hour := 120

	switch {
	case hour >= 5 && hour <= 10:
		fmt.Println("Pagi")
	case hour >= 11 && hour <= 14:
		fmt.Println("Siang")
	case hour >= 15 && hour <= 18:
		fmt.Println("Sore")
	case hour >= 19 && hour <= 23:
		fmt.Println("Malam")
	default:
		fmt.Println("Jam tidak valid")
	}
}
