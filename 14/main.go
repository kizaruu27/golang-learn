package main

import "fmt"

func main() {
	age := 20
	hasIDCard := true

	if age >= 18 && hasIDCard {
		fmt.Println("Boleh melakukan registrasi")
	} else if age >= 18 && !hasIDCard {
		fmt.Println("Harus memiliki KTP")
	} else {
		fmt.Println("Belum cukup umur")
	}
}
