package main

import "fmt"

var prev = -1
var curr = 1

func fibonacci() func() int {
	return func() int {
		sum := prev + curr
		prev = curr
		curr = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
