package publishers

import (
	"log"
	"os"

	"github.com/Khaled-Abdelal/job-crawler/crawler/worker"

	"github.com/streadway/amqp"
)

func PublishSearchWord(body []byte, ampqSession worker.AMPQSession) {
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
