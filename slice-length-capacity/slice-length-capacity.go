package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	//len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)
	// len=0 cap=6 []

	// Extend its length.
	s = s[:4]
	printSlice(s)
	// how come? len=4 cap=6 [2 3 5 7]

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
	// how come cap is 4? len=2 cap=4 [5 7]

	// extend beyond the capacity.
	//s = s[:8]
	//printSlice(s)
	// cause runtime error:  slice bounds out of range
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}