package crawlers

type Crawler interface {
	Crawl(jobTitle string) ([]Job, error)
}

type Job struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	URL         string `json:"URL"`
	Source      string `json:"source"`
	Description string `json:"description"`
	Location    string `json:"location"`
	CompanyName string `json:"companyName"`
}

func GetActiveCrawlers() []Crawler {
	result := []Crawler{
		NewIndeedCrawler(),
		NewRelocateMeCrawler(),
	}
	return result
}
