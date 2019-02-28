package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	fmt.Println("Started", s)
	start := time.Now()
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
	fmt.Println("Finished at ", time.Now().Sub(start), s)

}

func main() {
	s := []int{1, 2, 3, 4, -9, -5, 0, -2}
	c := make(chan int)

	go sum(s[len(s)/2:], c)
	go sum(s[:len(s)/2], c)

	x, y := <-c, <-c
	fmt.Printf("X is %v\nY is %v\nSum is %v\n", x, y, x+y)
}
