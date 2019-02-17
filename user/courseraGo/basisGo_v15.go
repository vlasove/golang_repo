package main

import "fmt"

type Person struct {
	Id     int
	Name   string
	Adress string
}

type Account struct {
	Id      int
	Name    string
	Cleaner func(string) string
	Owner   Person
	Person
}

func main() {
	var acc Account = Account{
		Id:     1,
		Name:   "Evgen",
		Person: Person{3, "Evgeny Vlasov", "Hui2"},
	}

	acc.Owner = Person{2, "Evgeny Vlasov", "Hui"}

	fmt.Printf("%#v\n", acc)
	fmt.Println(acc.Adress)

}
