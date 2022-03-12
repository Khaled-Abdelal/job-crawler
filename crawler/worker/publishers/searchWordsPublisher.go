package publishers

import (
	"context"
	"crawler/worker"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func PublishSearchWord(body []byte, ctx context.Context) {
	var ampqSession = worker.GetSessionFromContext(ctx)
	channelRabbitMQ := ampqSession.Channel
	err := channelRabbitMQ.Publish("", os.Getenv("SEARCH_WORD_TO_CRAWL_QUEUE"), false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})
	if err != nil {
		log.Printf("Error publishing message: %s", err)
	}
	log.Println("Sending a message to the channel")
}
