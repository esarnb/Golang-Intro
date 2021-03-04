package main

import "fmt"

func main() {
	// messages := make(chan string)

	// go func() {
	// 	messages <- "Nice Job"
	// }()
	// msg := <-messages
	// fmt.Println(msg)

	nums := []int{1, 2, 3, 4, 5}
	nums2 := []int{5, 11, 3, 2, 1}

	numChannel := make(chan int)
	go sum(nums, numChannel)
	go sum(nums2, numChannel)

	sum1, sum2 := <-numChannel, <-numChannel
	fmt.Println(sum1, sum2)
}

func sum(nums []int, c chan int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	c <- sum
}
