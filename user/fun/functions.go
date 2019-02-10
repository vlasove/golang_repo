package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Printf("The addition of %v and %v is equal to %v\n", 10, 15, add(10, 15))
}
