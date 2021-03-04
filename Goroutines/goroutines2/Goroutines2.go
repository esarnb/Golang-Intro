package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		countSheep()
	}()

	go func() {
		defer wg.Done()
		countFish()
	}()

	wg.Wait()
	fmt.Println("Cock")
}

func countFish() {
	defer fmt.Println("Fish done")

	for i := 1; i < 4; i++ {
		fmt.Printf("%d fish\n", i)
		time.Sleep(time.Second * 1)
	}
}

func countSheep() {
	defer fmt.Println("Sheep done")
	for i := 1; i < 4; i++ {
		fmt.Printf("%d sheep\n", i)
		time.Sleep(time.Second * 1)
	}
}
