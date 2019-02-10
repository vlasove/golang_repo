package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	x := "Hello"
	y := "Jacob"

	fmt.Printf("Original strings: %v and %v.\n", x, y)
	a, b := swap(x, y)
	fmt.Printf("Swapping strings: %v and %v.\n", a, b)

}
