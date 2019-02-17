package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	gypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(gypot(5, 12))
	fmt.Println("Compute works: ", compute(gypot))
	fmt.Println("Compute also works : ", compute(math.Pow))

}
