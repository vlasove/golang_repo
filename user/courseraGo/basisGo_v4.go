package main

import "fmt"

const pi = 3.1415

const (
	hello = "Hello"
	e     = 2.718
)

const (
	zero = iota
	one
	two
	three
	four
)

func main() {

	fmt.Println("Ready")
	fmt.Println(zero, one, two, three, four)

}
