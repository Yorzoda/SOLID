package main

import (
	"github.com/SOLID/SRP/jeffrey_example/salesReporter"
	"time"
)

func main() {
	start := time.Now().AddDate(0, 0, -10)
	end := time.Now()
	c := salesReporter.NewConsoleOutput()

	salesReporter.Report(start, end, c)

}
