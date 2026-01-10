package main

import "fmt"

func main() {
	names := [...]string{"Sumanto", "Rusdi", "Bambang", "Amba"}

	for _, name := range names {
		fmt.Println("Nama:", name)
		if name == "Bambang" {
			fmt.Printf("Ditemukan %s, looping dihentikan \n", name)
			break
		}
	}
}
