package main

import "fmt"

func Add(x, y float64) float64 {
	return x + y
}

func main() {
	fmt.Printf("The sum of %v and %v is %v\n", 2, 3, Add(2, 3))
}
