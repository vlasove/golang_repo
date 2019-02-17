package main

import "fmt"

func main() {
	fmt.Println("Ready")
	a := 2
	b := &a
	*b += 4
	fmt.Println(a)
	c := &a

	d := new(int)
	*d = 12
	fmt.Println(a)
	*c = *d
	fmt.Println(a)
	*d = 13
	fmt.Println(a)
}
