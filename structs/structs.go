package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	v.X = 4
	fmt.Println(v)

	p := &v
	p.X = 1e9
	fmt.Println(v)

	v2 := Vertex{X: 1}
	v3 := Vertex{}
	fmt.Println(v2)
	fmt.Println(v3)
}