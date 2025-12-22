package main

import "fmt"

func main() {
	myString1, myString2 := "Selamat", "Kerja"
	
	// Perbandingan >
	fmt.Printf("Perbandingan menggunakan \">\" : %t\n", myString1 > myString2)
	// Perbandingan <
	fmt.Printf("Perbandingan menggunakan \"<\" : %t\n", myString1 < myString2)
	// Perbandingan =
	fmt.Printf("Perbandingan menggunakan \"=\" : %t\n", myString1 == myString2)
}