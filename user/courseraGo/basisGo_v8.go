package main

import "fmt"

func main() {
	fmt.Println("Ready")
	var test [3]int
	var buf0 []int

	fmt.Printf("%T and %#v\n", test, test)
	fmt.Printf("%T and %#v\n", buf0, buf0)

	buf1 := []int{}
	buf2 := []int{42}

	buf3 := make([]int, 0)
	buf4 := make([]int, 5)
	buf5 := make([]int, 5, 10)

	fmt.Printf("%T and %#v\n", buf1, buf1)
	fmt.Printf("%T and %#v\n", buf2, buf2)
	fmt.Printf("%T and %#v\n", buf3, buf3)
	fmt.Printf("%T and %#v\n", buf4, buf4)
	fmt.Printf("%T and %#v\n", buf5, buf5)

	var buff []int
	//for i := 0; i < 100; i++ {
	//	buff = append(buff, i, i+1)
	//	fmt.Printf("%T with len = %v and cap= %v \n", buff, len(buff), cap(buff))
	//}
	otherSLice := make([]int, 3)
	buff = append(buff, otherSLice...)
	fmt.Println(buff)

	buff1 := []int{0, 1, 2, 3, 4, 5, 6}

	newBuff := buff1[:]

	fmt.Println(buff1, newBuff)
	newBuff[0] = 1000
	fmt.Println(buff1, newBuff)
	newBuff = append(newBuff, 10000, 10123)
	newBuff[0] = 999
	fmt.Println(buff1, newBuff)

	newBuff = make([]int, len(buff1), cap(buff1))
	copy(newBuff, buff1)
	fmt.Println(newBuff)
	fmt.Println(buff1)

}
