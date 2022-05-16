package main

import (
	"log"
	"os"

	"github.com/Khaled-Abdelal/job-crawler/keywords/cron"
	"github.com/Khaled-Abdelal/job-crawler/keywords/data"
	"github.com/Khaled-Abdelal/job-crawler/keywords/worker"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvFile()
	db, _ := data.GetDBConnection()
	data.SeedKeyWords("./data/seedKeyWords.csv", db) // seed the db with popular job key words to crawl
	ampqSession := worker.RabbitMQSetUp()            // holds the rabbitMQ session accross the app
	s := cron.RunSearchWordsCron(ampqSession, db)
	s.StartBlocking()
}

func loadEnvFile() {
	env := os.Getenv("APP_ENV")
	if env == "production" {
		return
	}
	if env == "" {
		env = "development"
	}
	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
