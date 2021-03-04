package main

import (
	"errors"
	"fmt"
)

// Strucure of functions
// func NameOfFunc(param1 type, param2 type) returnType {...}
// func NameOfFunc(param1 type, param2 type) (returnType1, returnType2) {...}

func main() {
	added := add(1, 2)
	fmt.Println(added)

	checked, err := check(1, 2)
	if checked {
		fmt.Println("Values are the same")
	} else {
		fmt.Println("Error: ", err)
	}
}

func add(a int, b int) int {
	return a + b
}

func check(a int, b int) (bool, error) {
	if a != b {
		return false, errors.New("values are not the same")
	}

	return true, nil
}
