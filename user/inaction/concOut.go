package main

import (
	"fmt"
	"time"
)

func count() {
	for i := 0; i < 7; i++ {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func main() {

	go count()
	time.Sleep(time.Second * 3)
	fmt.Println("Hello, John")
	time.Sleep(time.Second * 5)

}
