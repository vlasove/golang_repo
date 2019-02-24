package main

import (
	"fmt"

	flag "github.com/jessevdk/go-flags"
)

var opts struct {
	Name    string `short:"n" long:"name" default:"World" description:"A name to say hello to"`
	Spanish bool   `short:"s" long:"spanish" description:"Use Spanish language"`
}

func main() {
	_, err := flag.Parse(&opts)
	if err != nil {
		fmt.Println("can't parse some flags: ", err)
	}

	if opts.Spanish {
		fmt.Printf("Hola %s!\n", opts.Name)

	} else {
		fmt.Printf("Hello %s!\n", opts.Name)
	}
}
