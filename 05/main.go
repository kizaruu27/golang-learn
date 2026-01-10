package main

import "fmt"

func main() {
	scores := [5]int{90, 65, 70, 100, 80}
	min := 0
	max := 0

	for i, score := range scores {
		if i == 0 {
			min = score
			max = score

			continue
		}

		if score < min {
			min = score
		}

		if score > max {
			max = score
		}
	}

	fmt.Println("Nilai tertinggi:", max)
	fmt.Println("Nilai terendah:", min)
}
