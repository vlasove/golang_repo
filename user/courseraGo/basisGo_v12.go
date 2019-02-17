package main

import "fmt"

func doNothing() {
	fmt.Println("it's regular function")
}

func main() {
	fmt.Println("Ready")
	doNothing()

	//incognito dunctions
	func(in string) {
		fmt.Println("incognito function out is: ", in)
	}("nobody")

	printer := func(in string) {
		fmt.Println("printer out is : ", in)
	}
	printer("lol")
	printer("kek")
	printer("cheburek")

	type strFuncType func(string)

	worker := func(callback strFuncType) {
		callback("as callback")
	}

	worker(printer)

	//embedding
	prefixer := func(prefix string) strFuncType {
		return func(in string) {
			fmt.Printf("[%s] %s\n", prefix, in)
		}
	}

	succesLogger := prefixer("Success")
	succesLogger("expected behaviour")

}
