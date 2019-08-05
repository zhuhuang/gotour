package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	// Only the sender should close a channel, never the receiver.
	// Sending on a closed channel will cause a panic.
	// Another note: Channels aren't like files;
	// you don't usually need to close them.
	// Closing is only necessary when the receiver must be told there are no
	// more values coming, such as to terminate a range loop.
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	// The loop for i := range c receives values from the channel repeatedly until it is closed.
	for i := range c {
		fmt.Println(i)
	}
}
