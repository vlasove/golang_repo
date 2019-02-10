package main

import "fmt"

func main() {

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	b := [3]int{1, 2, 3}
	fmt.Println(b)

	items := [6]int{1, 2, 3, 4, 5, 6}

	var q []int = items[0:4]
	fmt.Println(q)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a1 := names[0:3]
	b1 := names[0:2]

	fmt.Println(a1, b1)
	b1[0] = "Derek"
	fmt.Println(a1, b1)

}
