package main

import "fmt"

func getSalary() int {
	return 5500000
}

func main() {
	if salary := getSalary(); salary >= 5000000 {
		fmt.Println("Gaji mencukupi")
	} else {
		fmt.Println("Gaji belum mencukupi")
	}
}
