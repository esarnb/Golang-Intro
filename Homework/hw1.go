package main

// Start Time: 7:56pm

/*		Pseudocode

= > Run program
= 	>Ask user what method to use:
=	>Add, Subtract, Multiply, Divide, Modulus, exit
=		> Run the respective function
=			> Remember error handling, [if no item, try again]
=			> Return useful value
=	> Loop until told to exit

*/

/*   Extra Notes

I realize i could change "isInArray" to be
a switch case as well for 1-5 choices,
but too lazy to delete tons of code and to
debug again so ima leave it like this

I also Removed useless functional methods,
aka func add() sub() mult() divide()
Because well, switch cases.

I only validated inputs, thats all i needed

*/

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to Esar's Calculator in Golang!")
	nextMethod := false

	for {
		choice := pickChoice()
		fmt.Println("You chose", choice)
		input1 := pickInputAB("first")
		fmt.Println("Input 1:", input1)
		input2 := pickInputAB("second")
		fmt.Println("Input 2:", input2)

		fmt.Println("\n= = = Output = = =")

		output := 0.0
		printMethod := ""
		switch choice {
		case "1":
			output = input1 + input2
			printMethod = "+"
			break
		case "2":
			output = input1 - input2
			printMethod = "-"
			break
		case "3":
			output = input1 * input2
			printMethod = "*"
			break
		case "4":
			output = input1 / input2
			printMethod = "/"
			break
		case "5":
			output = math.Mod(input1, input2)
			printMethod = "%"
			break
		default:
			checkExit("c")
		}

		fmt.Printf("%v %s %v = %v", input1, printMethod, input2, output)

		fmt.Println("\n= = = = = = = = = =")

		nextMethod = runAgain()
		if nextMethod != true {
			checkExit("c")
		}
	}
}

// pickChoice: Keeps asking user to pick a method until chosen.
// returned: Validated float64 typed number
func pickChoice() string {
	choice := ""
	choices := []string{"1", "2", "3", "4", "5", "c"}

	fmt.Println("What would you like to do?\n 1. Add\n 2. Subtract\n 3. Multiply\n 4. Divide\n 5. Modulus \n Press c to exit anytime in the program.")
	_, err1 := fmt.Scanln(&choice)

	for err1 != nil || isInArray(choice, choices) != true {
		fmt.Println("Please choose a valid integer choice (1-5) or c to exit.")
		_, err1 = fmt.Scanln(&choice)
	}

	checkExit(choice)

	return choice
}

// pickInputAB: Keeps asking user to input a float value and validates it.
// Param 1: Placeholder for which question is being served.
// returned: Validated float64 typed number
func pickInputAB(inputType string) float64 {
	//
	choice := ""
	var choiceFloat float64

	// do {...} while {...}
	for {

		// Ask as Question 1 or 2, then as error handler
		fmt.Println("Please choose a valid integer/decimal number for your " + inputType + " number: ")
		_, err := fmt.Scanln(&choice)

		// Input received?
		if err == nil {

			checkExit(choice)

			// Check input if its a valid float.
			choiceFloat, err = strconv.ParseFloat(choice, 64)
		}

		// If input is a valid float, return valid value.
		if err == nil {
			break
		}
	}

	return choiceFloat
}

/* https://stackoverflow.com/questions/15323767/does-go-have-if-x-in-construct-similar-to-python */
// isInArray: Checks if an element (string) is within an array (string).
// Param 1: The element to check with.
// Param 2: The array to check against.
// returned: True/False, whether or not the element was found.
func isInArray(a string, list []string) bool {

	// For each element (b) in the array list,
	for _, b := range list {

		// if b is finally found, stop checking and return 'found'.
		if b == a {
			return true
		}
	}

	return false
}

// checkExit: Checks if an input (string) is the special key to exit.
// Param 1: The user input to check against.
func checkExit(letter string) {
	if letter == "c" {
		fmt.Println("Goodbye!")
		os.Exit(0)
	}
}

// runAgain: Asks user to continue or quit; y for yes or n for no or c to exit
func runAgain() bool {
	choice := "n"
	fmt.Println("Run another method? [y/n]")
	_, err1 := fmt.Scanln(&choice)

	if err1 != nil {
		panic(err1)
	} else if choice == "y" {
		fmt.Println(".\n.\n.\n.\n.\n.")
		return true
	} else if choice == "c" || choice == "n" {
		checkExit("c")
	} else {
		return runAgain()
	}

	return false
}

// Ended: 9:55pm
