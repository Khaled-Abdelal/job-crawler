module github.com/Khaled-Abdelal/job-crawler/indexer

go 1.13

// replace github.com/Khaled-Abdelal/job-crawler/crawler/crawlers v0.0.0-local => ./../crawler

// replace github.com/Khaled-Abdelal/job-crawler/indexer/worker => ./worker

// replace github.com/Khaled-Abdelal/job-crawler/indexer/worker/consumers => ./worker/consumers

// replace github.com/Khaled-Abdelal/job-crawler/indexer/server => ./server

// replace github.com/Khaled-Abdelal/job-crawler/indexer/indexer => ./indexer

// replace github.com/Khaled-Abdelal/job-crawler/indexer/proto/jobservice => ./proto/jobservice

require (
	github.com/Khaled-Abdelal/job-crawler/crawler v0.0.0-20220515050728-cb1c1ae06493 // indirect
	//github.com/Khaled-Abdelal/job-crawler/crawler v0.0.0-local
	github.com/elastic/elastic-transport-go/v8 v8.1.0 // indirect
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20220324153036-fd6f0897d613
	github.com/isayme/go-amqp-reconnect v0.0.0-20210303120416-fc811b0bcda2
	github.com/joho/godotenv v1.4.0
	google.golang.org/grpc v1.32.0 // or newer
	google.golang.org/protobuf v1.28.0
)

//replace github.com/Khaled-Abdelal/job-crawler/crawler v0.0.0-local => ../crawler
