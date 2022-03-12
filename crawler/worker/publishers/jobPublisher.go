package publishers

import (
	"context"
	"crawler/crawlers"
	"crawler/worker"
	"encoding/json"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func PublishJobs(jobs []crawlers.Job, ctx context.Context) {
	var ampqSession = worker.GetSessionFromContext(ctx)
	channelRabbitMQ := ampqSession.Channel

	for _, job := range jobs {
		body, err := json.Marshal(job)
		if err != nil {
			log.Println(err, "Error encoding JSON")
		}
		err = channelRabbitMQ.Publish("", os.Getenv("CRAWLED_JOBS_QUEUE"), false, false, amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         body,
		})
		if err != nil {
			log.Printf("error %s publishing job %s", err, job)
		}
	}
}
