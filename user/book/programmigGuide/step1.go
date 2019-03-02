package main

import (
	"fmt"
	"time"
)

func Publish(text string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println(text)
	}()
}
func main() {

	Publish("LOlKek", 2*time.Second)
	time.Sleep(time.Second * 5)
	fmt.Println("5 seconds later")

}
