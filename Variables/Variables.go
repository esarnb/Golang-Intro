package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("hey")

	// First way, infer the type from the value
	// VarName := value
	myName := "Emp"
	// Second way, define the type
	// var VarName type = value
	var hisName string = "Other"

	fmt.Println(myName, hisName)

	// xd := 1.5
	xd := 1.5
	fmt.Println(xd)

	// Casting
	// Cast float to int (rounds down)
	xdAsInt := int(xd)
	fmt.Println(xdAsInt)
	// Strconv lib, parses strings into other vars

	strNum := "5"
	strInt, err := strconv.Atoi(strNum)
	// Catch error
	if err != nil {
		panic(err)
	}
	fmt.Println(strInt)
	// Ignore assignemnt to variable
	strInt2, _ := strconv.Atoi("1.5")
	fmt.Println(strInt2)
}
