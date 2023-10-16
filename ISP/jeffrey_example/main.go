package jeffrey_example

import "fmt"

type WorkableInterface interface {
	work()
}

type SleepableInterface interface {
	sleep()
}

type Capitan struct {
}

type Adroid struct {
}

func (c Capitan) work() {
	fmt.Println("Работа не волк работа work,волк(walk) это ходить)")
}

func (c Capitan) sleep() {
	fmt.Println("Capitan sleeper build")
}

func (a Adroid) work() {
	fmt.Println("Only work and nothing else!")
}
