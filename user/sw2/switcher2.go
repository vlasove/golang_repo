package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	default:
		fmt.Println("Too far away")

	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17 && t.Hour() > 12:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening!")
	}
}
