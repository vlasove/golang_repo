package person

import "fmt"

func NewPerson(id int, name, secret1 string) *Person {
	return &Person{
		Id:     id,
		Name:   name,
		secret: secret1,
	}
}

func GetSecret(p *Person) string {
	return p.secret
}

func PrintSecret(p *Person) {
	fmt.Println(p.secret)
}
