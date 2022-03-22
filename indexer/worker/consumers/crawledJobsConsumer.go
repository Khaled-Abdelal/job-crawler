package consumers

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/Khaled-Abdelal/job-crawler/indexer/worker"

	"github.com/Khaled-Abdelal/job-crawler/crawler/crawlers"
)

func CrawledJobsConsumer(ctx context.Context) {
	var ampqSession = worker.GetSessionFromContext(ctx)
	channelRabbitMQ := ampqSession.Channel

	messageChannel, err := channelRabbitMQ.Consume(
		os.Getenv("CRAWLED_JOBS_QUEUE"),
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range messageChannel {
			log.Printf("recived message on the channel %s", d.Body)
			job := &crawlers.Job{}
			err := json.Unmarshal(d.Body, job)
			if err != nil {
				panic(err)
			}

		}
	}()
	<-forever
}
