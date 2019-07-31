package main

import "fmt"

// If the backing array of s is too small to fit all the given values a bigger array will be allocated.
func main() {
	var s []int
	printSlice(s)
	// len=0 cap=0 []

	// append works on nil slices
	s = append(s, 0)
	printSlice(s)
	// len=1 cap=1 [0]

	// the slice grows as needed
	s = append(s, 1)
	printSlice(s)
	// len=2 cap=2 [0 1]

	// we can add more than one element at a time
	s = append(s, 2, 3, 4)
	printSlice(s)
	// how come cap is 6? len=5 cap=6 [0 1 2 3 4]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
