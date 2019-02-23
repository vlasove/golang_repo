package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Println(num)

	}
	fmt.Println("End of printCount")
}

func main() {
	c := make(chan int)
	a := []int{123, 123, 42, 14, 43, 2123, 53, 0, 91, 9, -1}
	go printCount(c)

	for _, val := range a {
		c <- val
		fmt.Println(val * 2)

	}

	time.Sleep(time.Second * 10)
	fmt.Println("End of main")
}
