package main

import "fmt"

func main() {
	// No. 11
	// Anonymus func
	greet := func(name string) string {
		return fmt.Sprintf("Halo %s, selamat datang!", name)
	}
	fmt.Println(greet("Bambang"))
}
