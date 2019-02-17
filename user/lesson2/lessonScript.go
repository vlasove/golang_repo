package main

import "fmt"

func panicExample(){
	fmt.Println("Hello")

	defer func () {
		if r:= recover(); r!= nil {
			fmt.Println("Lol %v", r)
		}
	}

	panic(nil)

	fmt.Println("End")
}

func main() {
	fmt.Println("Ready")
	//Объектная модель в Golang
	panicExample()
}
