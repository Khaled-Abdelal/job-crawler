package crawlers

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

type indeedCrawler struct{}

func NewIndeedCrawler() indeedCrawler {
	return indeedCrawler{}
}

func (indeedCrawler) Crawl(jobTitle string, ch chan Job) {
	cc := colly.NewCollector()
	cc.OnResponse(func(r *colly.Response) {
		log.Println("Done Visiting: ", r.StatusCode)
	})
	cc.OnRequest(func(r *colly.Request) {
		log.Println("Visiting: ", r.URL.String())
	})
	cc.OnHTML("li div[class^=job_]", func(e *colly.HTMLElement) {
		job := Job{}
		job.Title = e.ChildText("h2")
		job.URL = "https://www.indeed.com" + e.Attr("href")
		job.Source = "Indeed"
		job.Description = e.ChildText("li")
		job.Location = e.ChildText(".companyLocation")
		job.CompanyName = e.ChildText(".companyName")
		ch <- job
	})
	searchURL := fmt.Sprintf("https://www.indeed.com/q-%s-jobs.html", jobTitle)
	cc.Visit(searchURL)
	cc.Wait()
}
