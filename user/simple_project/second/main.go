package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func foo(i int) {
	fmt.Printf("The square root of %v is %v\n", i, math.Sqrt(float64(i)))
	time.Sleep(time.Second * 1)

}

func main() {
	fmt.Println("Ready!")
	fmt.Println(math.Sqrt(4.0))

	for i := 0; i < 20; i++ {

		go foo(i)

	}
	time.Sleep(time.Second * 5)

	fmt.Println("DOne")

	fmt.Printf("A number form 1 to 100 is %v", rand.Intn(100))
}
