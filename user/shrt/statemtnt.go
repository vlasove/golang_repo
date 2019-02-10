package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%v >= %v\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(11, 20, 15))
}
