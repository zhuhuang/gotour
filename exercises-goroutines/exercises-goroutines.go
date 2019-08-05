package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values from the tree to the channel ch
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Walker(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same determines whether the trees t1 and t2 contain the same values
func Same(t1, t2 *tree.Tree) bool {
	ct1 := make(chan int)
	ct2 := make(chan int)

	go Walker(t1, ct1)
	go Walker(t2, ct2)

	for {
		v1, ok1 := <-ct1
		v2, ok2 := <-ct2

		if ok1 && ok2 {
			if v1 != v2 {
				return false
			}
		} else if !ok1 && !ok2 {
			return true
		} else {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
