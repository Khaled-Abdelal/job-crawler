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

func (indeedCrawler) Crawl(jobTitle string) ([]Job, error) {
	var js []Job
	cc := colly.NewCollector()
	cc.OnResponse(func(r *colly.Response) {
		log.Println("Done Visiting: ", r.StatusCode)
	})
	cc.OnRequest(func(r *colly.Request) {
		log.Println("Visiting: ", r.URL.String())
	})
	cc.OnHTML("a[id^=job_]", func(e *colly.HTMLElement) {
		temp := Job{}
		temp.Title = e.ChildText("h2")
		temp.URL = "https://www.indeed.com" + e.Attr("href")
		temp.Source = "Indeed"
		temp.Description = e.ChildText("li")
		temp.Location = e.ChildText(".companyLocation")
		temp.CompanyName = e.ChildText(".companyName")
		js = append(js, temp)
	})
	searchURL := fmt.Sprintf("https://www.indeed.com/q-%s-jobs.html", jobTitle)
	cc.Visit(searchURL)
	cc.Wait()
	return js, nil
}
