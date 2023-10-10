package salesReporter

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type SalesRepo struct {
	ID        int
	charge    int
	createdAt time.Time
}

func NewSalesRepo() *SalesRepo {
	return &SalesRepo{}
}

func (s SalesRepo) GetTotalSales(db *gorm.DB, start, end time.Time) int {

	err := db.Model(SalesRepo{}).
		Select("sum(charge)").
		Where("createdAt BETWEEN  ? And ?", start, end).
		Scan(&s.charge).Error
	if err != nil {
		log.Println(err)
	}
	totalSales := s.charge / 100

	return totalSales
}
