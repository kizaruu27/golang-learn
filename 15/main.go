// No.15
// Higher order function

package main

import "fmt"

func processNumber(value int, operation []func(int) int) int {
	result := value
	for _, opt := range operation {
		result = opt(result)
	}

	return result
}

func main() {
	double := func(n int) int {
		return n * 2
	}

	addTen := func(n int) int {
		return n + 10
	}

	timesTen := func(n int) int {
		return n * 10
	}

	result := processNumber(5, []func(int) int{double, addTen, timesTen})
	fmt.Println(result)
}
