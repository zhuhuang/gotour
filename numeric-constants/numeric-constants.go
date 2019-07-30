package main

import "fmt"

const (
	Big = 1 << 100

	Small = Big >> 99
)

func needInt(x int) int {
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// fmt.Printf("Type of Big %T\n", Big) // overflow
	fmt.Printf("Type of Small %T\n", Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
