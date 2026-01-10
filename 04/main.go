package main

import "fmt"

func main() {
	myNumbers := [5]int{10, 4, 3, 7, 6}

	for i := 0; i < len(myNumbers); i++ {
		if myNumbers[i]%2 != 0 {
			continue
		}

		fmt.Printf("Angka genap: %d \n", myNumbers[i])
	}
}
