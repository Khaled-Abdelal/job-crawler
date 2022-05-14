package main

import (
	"log"
	"os"

	"github.com/Khaled-Abdelal/job-crawler/crawler/cron"
	"github.com/Khaled-Abdelal/job-crawler/crawler/data"
	"github.com/Khaled-Abdelal/job-crawler/crawler/worker"
	"github.com/Khaled-Abdelal/job-crawler/crawler/worker/consumers"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvFile()
	db, _ := data.GetDBConnection()
	data.SeedKeyWords("./data/seedKeyWords.csv", db)
	ampqSession := worker.RabbitMQSetUp() // holds the rabbitMQ session accross the app
	s := cron.RunSearchWordsCron(ampqSession, db)
	go s.StartBlocking()
	consumers.SearchWordsConsume(ampqSession)
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
