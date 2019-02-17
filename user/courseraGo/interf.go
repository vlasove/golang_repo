package main

import "fmt"

type Payer interface {
	Pay(int) error
}
type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("LOl bomzh")
	}
	w.Cash -= amount
	return nil
}

func Buy(p Payer) {
	err := p.Pay(10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("THanks for buying throu the %T\n\n", p)
}

func main() {
	fmt.Println("REady")
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)
}
