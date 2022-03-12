package cron

import (
	"context"
	"crawler/data"
	"crawler/worker/publishers"
	"encoding/json"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

func RunSearchWordsCron(ctx context.Context) gocron.Scheduler {
	s := gocron.NewScheduler(time.UTC)
	fn := func() { task(ctx) }
	s.Every(5).Seconds().Do(fn)
	return *s
}

func task(ctx context.Context) {
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
		publishers.PublishSearchWord(body, ctx)
	}
}
