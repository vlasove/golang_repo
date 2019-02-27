package main

import "fmt"

func main() {
	naturals := make(chan int)
	squared := make(chan int)

	go func() {
		defer close(naturals)
		for x := 0; x < 100; x++ {
			naturals <- x

		}
	}()

	go func() {
		defer close(squared)
		for x := range naturals {
			squared <- x * x
		}
	}()

	for x := range squared {
		fmt.Println(x)
	}

}
