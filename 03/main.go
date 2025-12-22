package main

import "fmt"

func main() {
	welcomeMessage := getWelcomeMessage()
	fmt.Println(welcomeMessage)
}

func getWelcomeMessage() string {
	return "Aku ingin menjadi programmer yang rajin ngoding!"
}