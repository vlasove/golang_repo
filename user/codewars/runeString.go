package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ready")
	sample := "Hello, World"
	temp := []rune{}
	sep := []rune("T")

	for _, v := range sample {
		if v != sep[0] {
			temp = append(temp, v)

		}

	}
	s := string(temp)
	fmt.Println(s)
}
