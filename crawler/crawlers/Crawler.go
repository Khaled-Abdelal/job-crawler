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

func (job *Job) validate() bool {
	if len(job.Title) == 0 || len(job.Description) == 0 || len(job.ID) == 0 {
		return false
	}
	return true
}

func GetActiveCrawlers() []Crawler {
	result := []Crawler{
		NewIndeedCrawler(),
		NewRelocateMeCrawler(),
	}
	return result
}
