package main

import "fmt"

func main() {
	myArr := []string{"a", "b", "c", "d", "e"}

	// Loops:
	// - No while loops in go
	// - for {code} same as while
	// Structure of for loop
	// Standard:
	// for i := 0; i < 10; i++ {code}
	// Range:
	// for index, value := range myArr {code}

	for i := 0; i < len(myArr); i++ {
		fmt.Printf("%d. %s\n", i+1, myArr[i])
	}

	for i, letter := range myArr {
		fmt.Printf("%d. %s \n", i+1, letter)
	}
}
