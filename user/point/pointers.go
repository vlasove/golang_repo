package main

import "fmt"

func main() {
	fmt.Println("Ready")
	i, j := 100, 100500
	p := &i
	*p = 200
	q := &j
	*q = *q * 100
	fmt.Printf("%T and %v\n", i, i)
	fmt.Printf("%T and %v\n", j, j)
	fmt.Printf("%T and %v\n", p, p)
	fmt.Printf("%T and %v\n", p, p)
	fmt.Println(i, j)
}
