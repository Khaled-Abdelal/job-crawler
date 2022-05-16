package cron

import (
	"log"
	"time"

	"github.com/Khaled-Abdelal/job-crawler/keywords/data"
	"github.com/Khaled-Abdelal/job-crawler/keywords/worker"
	"github.com/Khaled-Abdelal/job-crawler/keywords/worker/publishers"
	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
)

func RunSearchWordsCron(ampqSession worker.AMPQSession, db *gorm.DB) gocron.Scheduler {
	s := gocron.NewScheduler(time.UTC)
	fn := func() { task(ampqSession, db) }
	s.Every(3000).Seconds().Do(fn)
	return *s
}

func task(ampqSession worker.AMPQSession, db *gorm.DB) {
	log.Println("cron job activated")
	var sws []data.SearchWord
	db.Find(&sws)
	for _, sw := range sws {
		go publishers.PublishSearchWord(sw, ampqSession)
	}
}
