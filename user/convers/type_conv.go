package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 5, 7
	var z uint = uint(math.Sqrt(float64(x*x + y*y)))
	fmt.Printf("The hip of %v and %v is equal to %v\n", x, y, math.Sqrt(float64(x*x+y*y)))
	fmt.Printf("The z is equal to %v\n", z)
}
