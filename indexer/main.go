package main

import (
	"context"
	"log"
	"os"

	"github.com/Khaled-Abdelal/job-crawler/indexer/worker"
	"github.com/Khaled-Abdelal/job-crawler/indexer/worker/consumers"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvFile()
	ctx := context.Background()
	ctx = worker.RabbitMQSetUp(ctx)
	consumers.CrawledJobsConsumer(ctx)
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