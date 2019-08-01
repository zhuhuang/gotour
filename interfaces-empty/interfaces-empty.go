package main

import "fmt"

func main() {
	// An empty interface may hold values of any type. (Every type implements at least zero methods.)
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
