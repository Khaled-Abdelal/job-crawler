package data

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

func SeedKeyWords(filePath string, db *gorm.DB) error {
	lines, err := readData(filePath)
	if err != nil {
		log.Printf("Error seeding key words from csv file %s", err)
		return err
	}
	inserts := 0
	for _, line := range lines {
		for _, record := range line {
			sw := SearchWord{
				SearchWord: record,
			}
			result := db.Create(&sw) // pass pointer of data to Create
			if result.Error != nil {
				log.Printf("Error seeding key word %s from csv file %s", record, result.Error)
				continue
			}
			inserts = inserts + 1
		}
	}
	log.Printf("Seeded %d key words to DB", inserts)
	return nil
}

func readData(filePath string) ([][]string, error) {

	p, _ := filepath.Abs(filePath)
	f, err := os.Open(p)

	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
