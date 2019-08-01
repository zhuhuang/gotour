package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// pointer receiver
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// pointer parameter
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// value receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// value parameter
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// methods with pointer receivers take either a value or a pointer as the receiver when they are called

	v := Vertex{3, 4}
	v.Scale(2) // value is OK
	ScaleFunc(&v, 10) // must take pointer

	p := &Vertex{4, 3}
	p.Scale(3) // pointer is also fine
	ScaleFunc(p, 8) // must take pointer

	fmt.Println(v, p)

	// methods with value receivers take either a value or a pointer as the receiver when they are called

	fmt.Println(v.Abs()) // value
	fmt.Println(AbsFunc(v)) // must be value

	fmt.Println(p.Abs()) // pointer
	fmt.Println(AbsFunc(*p)) // must be value
}