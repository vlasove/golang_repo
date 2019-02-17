package main

import "fmt"

func main() {
	var a1 = [...]int{1, 2, 3}
	fmt.Println(a1)

	var a0 [3]int
	fmt.Println("a0 short", a0)
	fmt.Printf("a0 short %v\n", a0)
	fmt.Printf("a0 full %#v \n", a0)

}
