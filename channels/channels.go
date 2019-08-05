package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}

	// send to channel
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// create channel
	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	//receive from channel
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

}
