package main

import (
	"log"
	"os"

	"github.com/Khaled-Abdelal/job-crawler/crawler/worker"
	"github.com/Khaled-Abdelal/job-crawler/crawler/worker/consumers"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvFile()
	ampqSession := worker.RabbitMQSetUp() // holds the rabbitMQ session accross the app
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
