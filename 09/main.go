package main

import "fmt"

func main() {
	var stringPointer *string
	var intPointer = new(int)

	fmt.Printf("Address stringPointer : %v\n", stringPointer)
	fmt.Printf("Address intPointer : %v\n", intPointer)

	fmt.Printf("Value pointer: %v\n", *intPointer)
}