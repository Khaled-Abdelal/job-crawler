package indexer

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type contextKey string

const (
	ElasticSearchClientContextKey contextKey = "elasticSearchClient"
)

// Initialize a client with the default settings.
//
// An `ELASTICSEARCH_URL` environment variable will be used when exported.
//
func NewElasticSearchClient() *elasticsearch.Client {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	return es
}

func Index(client *elasticsearch.Client, index string, data string) error {
	res, err := client.Index(
		index,
		strings.NewReader(data),
	)
	if err != nil {
		log.Printf("error getting response: %s", err)
		return err
	}
	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%s", res.Status(), data)
		return errors.New("error indexing document")
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
			return err
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
	res.Body.Close()
	return nil
}

func _search(client *elasticsearch.Client, index string, query map[string]interface{}, from int, size int) (map[string]interface{}, error) {
	// parse the query
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	// Perform the search request.
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(index),
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithFrom(from),
		client.Search.WithSize(size),
		client.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	log.Println(strings.Repeat("=", 37))
	return r, nil
}
