package main

import "fmt"

func main() {
	level := 11

	switch {
	case level >= 10:
		fmt.Println("Level Tinggi")
		fallthrough
	case level >= 5:
		fmt.Println("Level Menengah")
	case level < 5:
		fmt.Println("Level Rendah")
	}
}
