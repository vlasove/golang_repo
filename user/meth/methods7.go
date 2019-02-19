package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:  %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := 1.0

	if x < 0 {
		err := ErrNegativeSqrt(x)
		return x, err
	}
	for {
		zPrevious := z
		z = z - (z*z-x)/(2*z)
		if math.Abs(zPrevious-z) < 0.000000000001 {
			break
		}
	}
	return z, nil
}

func main() {

	res, err := Sqrt(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	res, err = Sqrt(-2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}
