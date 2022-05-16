package publishers

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Khaled-Abdelal/job-crawler/keywords/data"
	"github.com/Khaled-Abdelal/job-crawler/keywords/worker"

	"github.com/streadway/amqp"
)

func PublishSearchWord(sw data.SearchWord, ampqSession worker.AMPQSession) error {
	body, err := json.Marshal(sw)
	if err != nil {
		log.Println(err, "Error encoding key word JSON")
		return err
	}
	channelRabbitMQ := ampqSession.Channel
	err = channelRabbitMQ.Publish("", os.Getenv("SEARCH_WORD_TO_CRAWL_QUEUE"), false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})
	if err != nil {
		log.Printf("Error publishing keyword: %s", err)
		return err
	}
	log.Printf("Sending keyword %s to the queue", sw.SearchWord)
	return nil
}
