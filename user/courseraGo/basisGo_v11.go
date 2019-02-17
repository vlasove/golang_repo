package main

import "fmt"

func singleIn(in int) int {
	return in * 10
}

func multiIn(a, b, c int) int {
	return a + b + c
}

func sumInt(in ...int) int {
	fmt.Printf("%T and %#v\n", in, in)
	var result int = 0
	for _, val := range in {
		result += val

	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(singleIn(15))
	fmt.Println(multiIn(10, 20, 30))
	fmt.Println(sumInt(nums...))
}
