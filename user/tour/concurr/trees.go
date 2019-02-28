package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walkRescue(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkRescue(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkRescue(t.Right, ch)
	}

}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkRescue(t, ch)
	close(ch)

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 {
			break
		}
		if v1 != v2 || ok1 != ok2 {
			return false
		}

	}
	return true
}

func main() {
	kek := tree.New(2)
	c := make(chan int)
	go Walk(kek, c)
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))

}
