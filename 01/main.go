package main

import "fmt"

func main() {
	age := 20
	if age >= 18 {
		fmt.Println("Boleh membuat KTP")
	} else {
		fmt.Println("Belum cukup umur")
	}
}
