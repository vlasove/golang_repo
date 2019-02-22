package main

import (
	"fmt"
	"time"
)

type MyError struct {
	What string
	When time.Time
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s ", e.When, e.What)

}

func run() error {
	return &MyError{"it didn't work", time.Now()}
}
func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
