package main

import (
	"fmt"
	"sync"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}

}

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)

}

func fibPrint(n int) {
	temp := fib(n)
	fmt.Printf("\n%v-th Fibonacci element is %v \n", n, temp)
}

func main() {
	go spinner(time.Millisecond * 50)
	//n := 45
	var wg sync.WaitGroup
	start := time.Now()

	for i := 4; i < 50; i += 5 {
		wg.Add(1)
		go func(t int) {
			fibPrint(t)
			wg.Done()
		}(i)

	}

	wg.Wait()

	fmt.Printf("The time duration for Fib calculus is : %v \n", time.Now().Sub(start))

}
