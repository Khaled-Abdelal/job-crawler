package crawlers

type Crawler interface {
	Crawl(jobTitle string) ([]Job, error)
}

type Job struct {
	title       string
	URL         string
	source      string
	description string
	location    string
	companyName string
}

func GetActiveCrawlers() []Crawler {
	result := []Crawler{}
	result = append(result, NewIndeedCrawler())

	return result
}
