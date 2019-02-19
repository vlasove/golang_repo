package main

import (
	"fmt"
	"math"
)

func makeRange(max int) []int {
	a := make([]int, max)
	for i := range a {
		a[i] = 1 + i
	}
	return a
}

func SuMin(m int) int64 {
	x := makeRange(m)
	y := makeRange(m)
	var sum int64 = 0

	for _, X := range x {
		for _, Y := range y {
			sum += int64(math.Min(float64(X), float64(Y)))
		}
	}
	return sum
}

func SuMax(m int) int64 {
	x := makeRange(m)
	y := makeRange(m)
	var sum int64 = 0

	for _, X := range x {
		for _, Y := range y {
			sum += int64(math.Max(float64(X), float64(Y)))
		}
	}
	return sum
}
func SumSum(m int) int64 {
	return SuMax(m) + SuMin(m)
}

func main() {
	fmt.Println(makeRange(10))
	temp := []int{6, 45, 999, 5000}
	for _, val := range temp {
		fmt.Printf("SuMin ---> %v\n", SuMin(val))
		fmt.Printf("SuMax ---> %v\n", SuMax(val))
		fmt.Printf("SumSum ---> %v\n", SumSum(val))
	}
}
