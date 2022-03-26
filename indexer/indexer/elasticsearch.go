package indexer

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
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
