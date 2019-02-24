package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to smbdy.")

var spanish, russian, ugandan bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language.")
	flag.BoolVar(&russian, "russian", false, "Use Russian Cyka Blyat language.")
	flag.BoolVar(&ugandan, "uganda", false, "Do u now de way?")
}

//func init() {
//	flag.BoolVar(&spanish, "spanish", false, "Use Spanish Language")
//	flag.BoolVar(&spanish, "s", false, "Use Spanish Language")

//}

func main() {
	flag.Parse()
	flag.PrintDefaults()
	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)

	}
	if russian == true {
		fmt.Printf("PRIVET BLYAT, TOVARISH %s!!!\n", *name)
	}

	if ugandan == true {
		fmt.Printf("U r owr kween, %s ?\n", *name)
	}
}
