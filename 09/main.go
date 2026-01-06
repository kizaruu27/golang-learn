package main

import "fmt"

func main() {
	isLoggedIn := true
	role := "admin"

	if isLoggedIn {
		if role == "admin" {
			fmt.Println("Selamat datang admin")
		} else {
			fmt.Println("Selamat datang user")
		}
	} else {
		fmt.Println("Silakan login terlebih dahulu")
	}
}
