package main

import (
	"fmt"
	"time"
)

func main() {
	// Define anonomyious function
	// It's just a function without a name, it's like JS
	// First way, put function into variable
	add := func(a int, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2))
	// Another way, define and call function outright
	func(a int, b int) {
		fmt.Println(a - b)
	}(1, 5)
	// Inside each function including main, you can use the 'defer' keyword
	// Like this: defer fmt.Println("xd")
	// defer will run what's after this ONCE the function finishes running
	// It doesn't matter where it's placed
	func() {
		defer func() {
			fmt.Println(1 + 5)
			fmt.Println("Func is done!")
		}()

		time.Sleep(1 * time.Second)
		fmt.Println("Counted 1 second")
	}()
}
