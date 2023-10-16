package my_example

import "fmt"

type HaveStunInterface interface {
	stun() bool
}

type HaveHexInterface interface {
	hex()
}

type WitchDockor struct {
}

type Lion struct {
}

func (wd WitchDockor) stun() bool {
	fmt.Println("stunned......")
	return true
}

func (l Lion) stun() bool {
	fmt.Println("mis......")
	return false
}

func (l Lion) hex() {
	fmt.Println("cluck cluck cluck(chicken sounds)")
}
