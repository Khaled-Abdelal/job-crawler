package main

import (
	"context"
	"log"
	"os"

	"github.com/Khaled-Abdelal/job-crawler/crawler/cron"
	"github.com/Khaled-Abdelal/job-crawler/crawler/worker"
	"github.com/Khaled-Abdelal/job-crawler/crawler/worker/consumers"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvFile()
	println("-----------------", os.Getenv("APP_ENV"))
	ctx := context.Background()
	ctx = worker.RabbitMQSetUp(ctx)
	s := cron.RunSearchWordsCron(ctx)
	go s.StartBlocking()
	consumers.SearchWordsConsume(ctx)
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
