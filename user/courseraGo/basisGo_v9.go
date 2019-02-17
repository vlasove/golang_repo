package main

import "fmt"

func main() {

	fmt.Println("Ready")
	//map learning

	var user map[string]string = map[string]string{
		"name":     "Pidor",
		"lastName": "Bolshoy",
	}
	fmt.Println(user["name"], user["lastName"])
	name, Error := user["kek"]
	if Error {
		fmt.Println("This Exist: ", name)

	} else {
		fmt.Println(Error)
	}
}
