package consumers

import (
	"context"
	"crawler/crawlers"
	"crawler/data"
	"crawler/worker"
	"crawler/worker/publishers"
	"encoding/json"
	"log"
	"os"
)

func SearchWordsConsume(ctx context.Context) {
	var ampqSession = worker.GetSessionFromContext(ctx)
	channelRabbitMQ := ampqSession.Channel

	messageChannel, err := channelRabbitMQ.Consume(
		os.Getenv("SEARCH_WORD_TO_CRAWL_QUEUE"),
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
			keywordDB := &data.SearchWord{}
			err := json.Unmarshal(d.Body, keywordDB)
			if err != nil {
				panic(err)
			}
			activeCrawlers := crawlers.GetActiveCrawlers()
			jobsCrawled := []crawlers.Job{}
			for _, c := range activeCrawlers {
				js, err := c.Crawl(keywordDB.SearchWord)
				if err != nil {
					log.Printf("error crawling jobs for crawler %s", c)
				}
				jobsCrawled = append(jobsCrawled, js...)
			}
			publishers.PublishJobs(jobsCrawled, ctx)
		}
	}()
	<-forever
}
