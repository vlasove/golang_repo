package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

func (p Person) UpdateName(name string) {
	p.Name = name
}

func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

type MySlice []int

func (sl *MySlice) Add(i int) {
	*sl = append(*sl, i)
}

func (sl *MySlice) Count() int {
	return len(*sl)
}

func main() {

	obj := Person{1, "Kek"}
	fmt.Println(obj.Name)
	obj.SetName("Nicky Minaje")
	fmt.Println(obj.Name)

	newAcc := Account{
		Id:   10,
		Name: "Pidor",
		Person: Person{
			Id:   11,
			Name: "Vaska",
		},
	}

	newAcc.SetName("Pidor2")
	fmt.Println(newAcc.Name, newAcc.Person.Name)

	sl := MySlice([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(sl, "Len is ", sl.Count())
	sl.Add(10)
	fmt.Println(sl, "Len is ", sl.Count())
}
