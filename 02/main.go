package main

import "fmt"

func main() {
	printGreeting("Samsudin")
}

func printGreeting(name string) {
	fmt.Printf("Halo, %s! Selamat belajar Golang \n", name)
}