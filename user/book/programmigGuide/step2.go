package main

import (
	"fmt"
	"time"
)

func Publish(text string, delay time.Duration) (wait <-chan struct{}) {

	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println(text)
		close(ch)

	}()
	return ch
}

func main() {
	wait := Publish("LOlKek", time.Second*2)
	<-wait

}
