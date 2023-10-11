package main

import "fmt"

type PaymentMethod interface {
	pay()
}

type Payment struct {
}

func (p Payment) Process(pm PaymentMethod) {
	pm.pay()
}

type DebitCard struct {
	PAN string
}

func (d DebitCard) pay() {
	///Transaction logic
	fmt.Print("Transaction succeeded for ", d.PAN)
}

func main() {
	p := Payment{}
	card := DebitCard{PAN: "123456789"}
	p.Process(card)
}
