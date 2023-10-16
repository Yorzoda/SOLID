package my_example

import "fmt"

type key interface {
	open() bool
}

type electronicKey struct {
}

type physicalKey struct {
}

func (e electronicKey) open() bool {
	fmt.Print("Put to open door")
	return true
}

func (p physicalKey) open() bool {
	fmt.Print("Insert and turn clockwise to open door")
	return true
}
