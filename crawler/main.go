package main

import (
	"context"
	"crawler/cron"
	"crawler/worker"
	"crawler/worker/consumers"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	ctx = worker.RabbitMQSetUp(ctx)
	s := cron.RunSearchWordsCron(ctx)
	go s.StartBlocking()
	consumers.SearchWordsConsume(ctx)
}
