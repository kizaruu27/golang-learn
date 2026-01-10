package main

import "fmt"

func main() {
	nilai := 80

	if nilai < 0 || nilai > 100 {
		fmt.Println("Nilai tidak valid")
	} else if nilai >= 0 && nilai <= 59 {
		fmt.Println("Tidak Lulus")
	} else {
		fmt.Println("Lulus")
	}
}
