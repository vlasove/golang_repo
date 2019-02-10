package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	i := 0
	for {
		z_previous := z
		z = z - (z*z-x)/(2*z)
		fmt.Printf("Current iteration is %v.\nCurrent value is %v.\n", i, z)

		if math.Abs(z_previous-z) < 0.00001 {
			break
		}
		i += 1

	}
	return z
}

func main() {
	fmt.Println(Sqrt(102409313))
}
