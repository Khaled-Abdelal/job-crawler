package data

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SearchWord struct {
	ID         uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	SearchWord string
}

func GetDBConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=password dbname=crawler port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&SearchWord{})
	db.Create(&SearchWord{SearchWord: "f developer"})
	return db, err
}
