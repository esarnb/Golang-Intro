package main

import "fmt"

func main() {
	// User input example:
	userName := ""
	userAge := 0
	// Ask for input
	fmt.Println("What is your name?")
	// Get input
	fmt.Scanln(&userName)

	fmt.Println("What is your age?")
	fmt.Scanln(&userAge)

	fmt.Printf("Your name is: %s, and you are %d old!", userName, userAge)
}
