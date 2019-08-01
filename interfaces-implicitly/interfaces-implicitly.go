package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// A type implements an interface by implementing its methods.
// No "implements" keyword like other languages, e.g., Java
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}