package crawlers

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gocolly/colly/v2"
)

type indeedCrawler struct{}

func NewIndeedCrawler() indeedCrawler {
	return indeedCrawler{}
}

func (indeedCrawler) Crawl(jobTitle string) ([]Job, error) {
	cc := colly.NewCollector()
	cc.Limit(&colly.LimitRule{
		DomainGlob:  "*indeed.*",
		Parallelism: 1,
		Delay:       5 * time.Second,
	})
	cc.OnResponse(func(r *colly.Response) {
		log.Println("Done Visiting: ", r.StatusCode)
	})
	cc.OnRequest(func(r *colly.Request) {
		log.Println("Visiting: ", r.URL.String())
	})
	var jobs []Job
	cc.OnHTML("li div[class^=job_]", func(e *colly.HTMLElement) {
		job := Job{}
		job.ID = e.ChildAttr("a[id^=job_]", "id")
		job.Title = e.ChildText("h2")
		job.URL = "https://www.indeed.com" + e.ChildAttr("a[id^=job_]", "href")
		job.Source = "Indeed"
		job.Description = e.ChildText("li")
		job.Location = e.ChildText(".companyLocation")
		job.CompanyName = e.ChildText(".companyName")
		if !job.validate() {
			log.Print("Error: got invalid job from indeed crawler for keyword %", jobTitle)
			return
		}
		jobs = append(jobs, job)
	})
	for i := 0; i < 5; i++ { // scrap 5 pages
		searchURL := fmt.Sprintf("https://www.indeed.com/jobs?q=%s&start=%d", url.QueryEscape(jobTitle), i*10)
		cc.Visit(searchURL)
	}
	cc.Wait()
	return jobs, nil
}
