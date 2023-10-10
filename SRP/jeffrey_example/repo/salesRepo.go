package repo

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type SalesRepo struct {
	ID        int
	charge    int
	createdAt time.Duration
}

func GetTotalSales(db *gorm.DB, start, end time.Time) int {
	var charge int

	err := db.Model(SalesRepo{}).
		Select("sum(charge)").
		Where("createdAt BETWEEN  ? And ?", start, end).
		Scan(&charge).Error
	if err != nil {
		log.Println(err)
	}
	totalSales := charge / 100

	return totalSales
}
