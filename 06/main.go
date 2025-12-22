package main

import "fmt"

func main() {
	// Int variable
	var myIntValue int = 50
	fmt.Printf("Value before: %d \n", myIntValue)
	
	// Pointer variable
	var myPointerInt *int = &myIntValue
	
	// Change the value via pointer
	*myPointerInt = 100
	
	// Output
	fmt.Printf("Value after: %d \n", myIntValue)
}