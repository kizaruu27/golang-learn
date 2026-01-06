package main

import "fmt"

func main() {
	// No.12
	// Immediately Invoked Func
	func(name string, age int, city string) {
		fmt.Println("Nama:", name)
		fmt.Println("Umur:", age)
		fmt.Println("Kota:", city)
	}("Bambang", 25, "Bogor")
}
