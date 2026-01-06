// No.13
// Function + variadic + return

package main

import "fmt"

func calculateAverage(values ...int) float64 {
	if len(values) == 0 {
		return 0
	} else {
		value := 0
		totalValues := len(values)

		for _, n := range values {
			value += n
		}
		average := float64(value) / float64(totalValues)
		return float64(average)
	}
}

func main() {
	avg := calculateAverage(10, 10, 30)
	fmt.Println("Rata-rata:", avg)
}
