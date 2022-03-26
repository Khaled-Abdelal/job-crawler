package main

import (
	"log"
	"os"

	"github.com/Khaled-Abdelal/job-crawler/indexer/indexer"
	"github.com/Khaled-Abdelal/job-crawler/indexer/worker"
	"github.com/Khaled-Abdelal/job-crawler/indexer/worker/consumers"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvFile()
	esClient := indexer.NewElasticSearchClient()
	rabbitMQSession := worker.RabbitMQSetUp() // holds the rabbitMQ Connection, channel
	consumers.CrawledJobsConsumer(rabbitMQSession, esClient)

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
