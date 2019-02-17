package main

import "fmt"

func main() {

	defer func() {
		fmt.Println("IN PANIC BLYAAAAT")
		if err := recover(); err != nil {
			fmt.Println("panic happend")
		}
	}()

	fmt.Println("Useful work")
	panic("BLYYYYYYYYYYAT")

}
