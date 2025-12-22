package main

import "fmt"

func main() {
	value1, value2, value3 := 100, 20, 50
	value1 += 10
	value2 += 50
	value3 += 60

	fmt.Println("+= pada value1 ", value1)
	fmt.Println("+= pada value2 ", value2)
	fmt.Println("+= pada value3 ", value3)


	value1 *= 30
	value2 *= 70
	value3 *= 40
	
	fmt.Println("*= pada value1 ", value1)
	fmt.Println("*= pada value2 ", value2)
	fmt.Println("*= pada value3 ", value3)

	value1 %= 7
	value2 %= 3
	value3 %= 9

	fmt.Println("%= pada value1 ", value1)
	fmt.Println("%= pada value2 ", value2)
	fmt.Println("%= pada value3 ", value3)
}