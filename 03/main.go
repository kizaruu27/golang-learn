package main

import "fmt"

func main() {
	score := 65

	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else if score >= 60 {
		fmt.Println("Grade D")
	} else {
		fmt.Println("Grade E")
	}
}
