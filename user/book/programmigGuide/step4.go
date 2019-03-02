package main

import "fmt"

func Generator() chan int {
	ch := make(chan int)
	go func() {
		n := 0
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}

	}()
	return ch
}

func main() {
	fmt.Println("Ready")

	num := Generator()
	for i := 0; i < 500; i++ {
		fmt.Println(<-num)

	}
	close(num)
}
