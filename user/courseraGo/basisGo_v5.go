package main

import "fmt"

type userId int

func main() {
	fmt.Println("Ready")
	idx := 1
	var usr userId = 42

	myUI := userId(idx)

	fmt.Printf("%T and %v\n", usr, usr)
	fmt.Printf("%T and %v\n", myUI, myUI)
}
