package main

import "fmt"

func main() {
	username := "admin"
	password := "12345"

	if username == "admin" && password == "12345" {
		fmt.Println("Login berhasil")
	} else {
		fmt.Println("Login gagal")
	}
}
