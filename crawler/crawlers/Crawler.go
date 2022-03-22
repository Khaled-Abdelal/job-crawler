package crawlers

type Crawler interface {
	Crawl(jobTitle string) ([]Job, error)
}

type Job struct {
	Title       string
	URL         string
	Source      string
	Description string
	Location    string
	CompanyName string
}

func GetActiveCrawlers() []Crawler {
	result := []Crawler{}
	result = append(result, NewIndeedCrawler())

	return result
}
