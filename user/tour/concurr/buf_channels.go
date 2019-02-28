package main

import "fmt"

func main() {
	c := make(chan int, 3)
	c <- 3
	c <- 2

	for i := 0; i < 2; i++ {
		fmt.Println(<-c)
	}
}
