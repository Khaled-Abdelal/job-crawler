package consumers

import (
	"encoding/json"
	"log"
	"os"
	"sync"

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
		log.Print(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range messageChannel {
			log.Printf("received message on the channel %s", d.Body)
			keywordDB := &data.SearchWord{}
			err := json.Unmarshal(d.Body, keywordDB)
			if err != nil {
				log.Print(err)
			}
			activeCrawlers := crawlers.GetActiveCrawlers()
			var wg sync.WaitGroup
			for _, c := range activeCrawlers {
				wg.Add(1)
				go func(c crawlers.Crawler, wg *sync.WaitGroup) {
					jobs, err := c.Crawl(keywordDB.SearchWord)
					if err != nil {
						log.Printf("Error crawler jobs for keyword: %s", keywordDB.SearchWord)
						return
					}
					err = publishers.PublishJobs(jobs, ampqSession) // pass jobs to be published for indexing
					if err != nil {
						log.Printf("Error Publishing jobs for keyword: %s", keywordDB.SearchWord)
						return
					}
					wg.Done()
				}(c, &wg)
			}
			wg.Wait()
		}
	}()
	<-forever
}
