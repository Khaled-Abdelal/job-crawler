package cron

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Khaled-Abdelal/job-crawler/crawler/data"
	"github.com/Khaled-Abdelal/job-crawler/crawler/worker"
	"github.com/Khaled-Abdelal/job-crawler/crawler/worker/publishers"
	"github.com/go-co-op/gocron"
)

func RunSearchWordsCron(ampqSession worker.AMPQSession) gocron.Scheduler {
	s := gocron.NewScheduler(time.UTC)
	fn := func() { task(ampqSession) }
	s.Every(5).Seconds().Do(fn)
	return *s
}

func task(ampqSession worker.AMPQSession) {
	log.Println("cron job activated")
	now := time.Now()
	sixHoursAgo := now.Add(time.Duration(-6) * time.Hour)
	db, _ := data.GetDBConnection()
	var sws []data.SearchWord
	db.Find(&sws, "updated_at < ?", sixHoursAgo)
	if len(sws) > 0 {
		firstSw := sws[0]
		body, err := json.Marshal(firstSw)
		if err != nil {
			log.Println(err, "Error encoding JSON")
		}
		publishers.PublishSearchWord(body, ampqSession)
	}
}
