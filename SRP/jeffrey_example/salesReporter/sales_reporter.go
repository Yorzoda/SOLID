package salesReporter

import "time"

func Report(start, end time.Time, output SalesOutputInterface) {
	salesRepository := NewSalesRepo()
	sales := salesRepository.GetTotalSales(nil, start, end)
	output.output(sales)
}
