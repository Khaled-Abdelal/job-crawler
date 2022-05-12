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
			ch := make(chan crawlers.Job, len(activeCrawlers)*100) // assume a buffer of 100 job per crawler
			var wg sync.WaitGroup
			for _, c := range activeCrawlers {
				wg.Add(1)
				go func(c crawlers.Crawler, ch chan crawlers.Job, wg *sync.WaitGroup) {
					c.Crawl(keywordDB.SearchWord, ch)
					wg.Done()
				}(c, ch, &wg)
			}
			wg.Wait()
			close(ch)                               // close the channel after all the sites has been crawled
			publishers.PublishJobs(ch, ampqSession) // pass channel with all received jobs to be published for indexing
		}
	}()
	<-forever
}
