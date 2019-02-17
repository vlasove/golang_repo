package main

import (
	"fmt"

	"github.com/user/courseraGo/pkg_ex/person"
)

func main() {

	p := person.NewPerson(1, "vevgeny", "secret_blya")

	//fmt.Printf("main.Person : %+v \n", p.secret)

	secret := person.GetSecret(p)
	fmt.Println("GetSectet: ", secret)
}
