package main

import (
	"time"

	"github.com/SOLID/SRP/jeffrey_example/repo"
)

func main() {
	start := time.Now().AddDate(0, 0, 10)
	end := time.Now()

	repo.GetTotalSales(nil, start, end)

}
