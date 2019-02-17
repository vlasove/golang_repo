package main

import "fmt"

func main() {

	fmt.Println("Ready")

	var (
		i         int     = 10
		autoInt           = 15
		bigInt    int64   = 1<<32 - 1
		unsInt    uint    = 100500
		unsBigInt uint64  = 1<<64 - 1
		pi        float32 = 3.1415
		e         float64 = 2.7
		kek       float64

		b     bool
		isOk  = true
		notOk = false
	)

	fmt.Printf("%T and %v\n", i, i)
	fmt.Printf("%T and %v\n", autoInt, autoInt)
	fmt.Printf("%T and %v\n", bigInt, bigInt)
	fmt.Printf("%T and %v\n", unsInt, unsInt)
	fmt.Printf("%T and %v\n", unsBigInt, unsBigInt)

	fmt.Println("============float part=============")
	fmt.Printf("%T and %v\n", pi, pi)
	fmt.Printf("%T and %v\n", e, e)
	fmt.Printf("%T and %v\n", kek, kek)

	fmt.Println("================bool part ==========")
	fmt.Printf("%T and %v\n", b, b)
	fmt.Printf("%T and %v\n", isOk, isOk)
	fmt.Printf("%T and %v\n", notOk, notOk)

}
