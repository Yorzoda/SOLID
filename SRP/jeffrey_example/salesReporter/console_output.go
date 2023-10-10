package salesReporter

import "fmt"

type ConsoleOutput struct {
}

func NewConsoleOutput() *ConsoleOutput {
	return &ConsoleOutput{}
}

func (c ConsoleOutput) output(sales int) {
	fmt.Printf("Sales:%v", sales)
}
