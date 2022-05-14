package indexer

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/Khaled-Abdelal/job-crawler/crawler/crawlers"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type GetJobsResponse struct {
	Jobs  []crawlers.Job
	Total int32
}

func GetJobs(client *elasticsearch.Client, term string, from int32, size int32) (GetJobsResponse, error) {
	searchFields := [2]string{"title^10", "description"}
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  term,
				"fields": searchFields,
			},
		},
	}
	res, err := _search(client, "my-index", query, from, size)
	if err != nil {
		log.Printf("Elastic search error: %s", err)
		return GetJobsResponse{}, nil
	}
	var jobs []crawlers.Job
	for _, hit := range res["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
		jobJson, err := json.Marshal(hit.(map[string]interface{})["_source"])
		if err != nil {
			log.Printf("Error parsing ElasticSearch result to json: %s", err)
			return GetJobsResponse{}, nil
		}
		var job crawlers.Job
		err = json.Unmarshal([]byte(jobJson), &job)
		if err != nil {
			log.Printf("Error parsing Searched Jobs: %s", err)
			return GetJobsResponse{}, nil
		}
		jobs = append(jobs, job)
	}
	total := int32(res["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	response := GetJobsResponse{
		Jobs:  jobs,
		Total: total,
	}
	return response, err
}

func IndexJobs(client *elasticsearch.Client, job crawlers.Job) error {
	jobBytes, err := json.Marshal(job)
	if err != nil {
		log.Printf("error parsing crawled job: %s", err)
		return err
	}
	req := esapi.IndexRequest{
		Index:      "my-index",
		DocumentID: job.ID,
		Body:       bytes.NewReader(jobBytes),
		Refresh:    "true",
	}
	return _index(client, req)
}
