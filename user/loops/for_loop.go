package main

import (
	"fmt"
	"time"
)

func loopInTIme(cycle int) {
	sum := 0
	start := time.Now()
	for i := 0; i < cycle; i++ {
		sum += i

	}
	finish := time.Now()
	fmt.Println(sum)
	fmt.Printf("Ended in %v \n", finish.Sub(start))

}

func main() {
	loopInTIme(10)
	loopInTIme(100)
	loopInTIme(100000000)
	sum := 0
	for sum < 1000 {
		sum += 20
	}
	fmt.Println(sum)
}
