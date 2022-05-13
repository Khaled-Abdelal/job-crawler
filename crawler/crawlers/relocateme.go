package crawlers

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type relocateMeCrawler struct{}

func NewRelocateMeCrawler() relocateMeCrawler {
	return relocateMeCrawler{}
}

func (relocateMeCrawler) Crawl(jobTitle string) ([]Job, error) {
	cc := colly.NewCollector()
	cc.Limit(&colly.LimitRule{
		DomainGlob:  "*relocate.me*",
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
	cc.OnHTML("div.jobs-list__job", func(e *colly.HTMLElement) {
		job := Job{}
		job.ID = strings.ReplaceAll(e.ChildAttr("div.job__title a", "href"), "/", "") // job id replace / to jot cause url parsing issues
		job.Title = e.ChildText("div.job__title b")
		job.URL = "https://relocate.me" + e.ChildAttr("div.job__title a", "href")
		job.Source = "Relocate.me"
		job.Description = e.ChildText("p.job__preview")
		job.Location = strings.ReplaceAll(e.ChildText("div.job__title"), job.Title, "")
		job.CompanyName = e.ChildText("div.job__company")
		if !job.validate() {
			log.Print("Error: got invalid job from relocate me crawler for keyword %", jobTitle)
			return
		}
		jobs = append(jobs, job)
	})
	for i := 0; i < 5; i++ { // scrap 5 pages
		searchURL := fmt.Sprintf("https://relocate.me/search?query=%s&page=%d", url.QueryEscape(jobTitle), i+1)
		cc.Visit(searchURL)
	}
	cc.Wait()
	return jobs, nil
}
