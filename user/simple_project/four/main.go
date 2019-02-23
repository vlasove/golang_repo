package main

import "fmt"

func main() {
	x := 15
	a := &x
	fmt.Println(*a)
	*a = 500
	fmt.Println(x)
}
