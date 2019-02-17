package main

import "fmt"

func getSomeVars() string {
	fmt.Println("Function Worked")
	return "getSOmeVars() finished"
}

func main() {
	defer fmt.Println(getSomeVars())
	fmt.Println("Work start")
}
