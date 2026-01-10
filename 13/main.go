package main

import "fmt"

func main() {
	isMember := true
	totalBelanja := 750000

	if isMember && totalBelanja >= 500000 {
		fmt.Println("Diskon 20%")
	} else if !isMember && totalBelanja >= 500000 {
		fmt.Println("Diskon 10%")
	} else {
		fmt.Println("Tidak dapat diskon")
	}
}
