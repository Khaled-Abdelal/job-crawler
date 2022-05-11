package consumers

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Khaled-Abdelal/job-crawler/crawler/crawlers"
	"github.com/Khaled-Abdelal/job-crawler/crawler/data"
	"github.com/Khaled-Abdelal/job-crawler/crawler/worker"
	"github.com/Khaled-Abdelal/job-crawler/crawler/worker/publishers"
)

func SearchWordsConsume(ampqSession worker.AMPQSession) {
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
			log.Printf("received message on the channel %s", d.Body)
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
			publishers.PublishJobs(jobsCrawled, ampqSession)
		}
	}()
	<-forever
}
