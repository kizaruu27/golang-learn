// No.14
// Function di dalam function

package main

import "fmt"

func counter() func() int {
	value := 0
	return func() int {
		value++
		return value
	}
}

func main() {
	count := counter()
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
}
