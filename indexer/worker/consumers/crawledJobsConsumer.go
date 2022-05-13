package consumers

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Khaled-Abdelal/job-crawler/indexer/indexer"
	"github.com/Khaled-Abdelal/job-crawler/indexer/worker"
	"github.com/elastic/go-elasticsearch/v8"

	"github.com/Khaled-Abdelal/job-crawler/crawler/crawlers"
)

func CrawledJobsConsumer(rabbitMQSession *worker.AMPQSession, elasticSearchClient *elasticsearch.Client) {
	channelRabbitMQ := rabbitMQSession.Channel

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
			log.Printf("received message on the channel %s", d.Body)
			job := &crawlers.Job{}
			err := json.Unmarshal(d.Body, job)
			if err != nil {
				log.Printf("error parsing crawled job: %s", err)
				return
			}
			err = indexer.IndexJobs(elasticSearchClient, *job)
			if err != nil {
				log.Printf("error parsing indexing crawled job: %s", err)
				return
			}
		}
	}()
	<-forever
}
