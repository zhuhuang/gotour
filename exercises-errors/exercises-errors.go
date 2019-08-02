package main

import (
	"fmt"
	"math"
)

const delta = 0.0001

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// return fmt.Sprintf("cannot Sqrt negative number: %v", e)
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	for {
		oldZ := z
		z -= (z * z - x) / (2 * z)
		if (math.Abs(oldZ - z) < delta) {
			break
		}
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

