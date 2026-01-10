package main

import "fmt"

func main() {
	year := 2020
	checkKabisat := func(year int) bool {
		if year%400 == 0 {
			return true
		} else if year%4 == 0 && year%100 != 0 {
			return true
		} else {
			return false
		}
	}

	if checkKabisat(year) {
		fmt.Println("Tahun kabisat")
	} else {
		fmt.Println("Bukan tahun kabisat")
	}
}
