package main

import (
	"fmt"
	"time"
)

func main() {
	go countFish()
	countSheep()
}

func countFish() {
	defer fmt.Println("Fish done")

	for i := 1; i < 10; i++ {
		fmt.Printf("%d fish\n", i)
		time.Sleep(time.Second * 1)
	}
}

func countSheep() {
	defer fmt.Println("Sheep done")
	for i := 1; i < 10; i++ {
		fmt.Printf("%d sheep\n", i)
		time.Sleep(time.Second * 1)
	}
}
