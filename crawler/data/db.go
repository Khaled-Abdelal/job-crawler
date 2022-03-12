package data

import (
	"os"
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
	dbConnectionString := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dbConnectionString), &gorm.Config{})
	db.AutoMigrate(&SearchWord{})
	db.Create(&SearchWord{SearchWord: "f developer"})
	return db, err
}
