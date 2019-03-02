package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {

		n := 0
		n++
		ch <- n
		close(ch)
	}()

	n := <-ch
	n++
	fmt.Println(n)
}
