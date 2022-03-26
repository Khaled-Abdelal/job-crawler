package consumers

import (
	"encoding/json"
	"log"
	"os"
	"strings"

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
				panic(err)
			}
			res, err := elasticSearchClient.Index(
				"my-index",
				strings.NewReader(string(d.Body)),
			)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%s", res.Status(), job.Title)
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
			res.Body.Close()
		}
	}()
	<-forever
}
