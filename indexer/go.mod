module github.com/Khaled-Abdelal/job-crawler/indexer

go 1.13

replace github.com/Khaled-Abdelal/job-crawler/crawler/crawlers => ./../crawler/crawlers

replace github.com/Khaled-Abdelal/job-crawler/indexer/worker => ./worker

replace github.com/Khaled-Abdelal/job-crawler/indexer/worker/consumers => ./worker/consumers

replace github.com/Khaled-Abdelal/job-crawler/indexer/server => ./server

replace github.com/Khaled-Abdelal/job-crawler/indexer/indexer => ./indexer

replace github.com/Khaled-Abdelal/job-crawler/indexer/proto/jobservice => ./proto/jobservice

require (
	github.com/Khaled-Abdelal/job-crawler/crawler v0.0.0-20220322201201-a8d98c540e9e
	github.com/PuerkitoBio/goquery v1.8.0 // indirect
	github.com/antchfx/htmlquery v1.2.4 // indirect
	github.com/antchfx/xmlquery v1.3.10 // indirect
	github.com/elastic/elastic-transport-go/v8 v8.1.0 // indirect
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20220324153036-fd6f0897d613
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/isayme/go-amqp-reconnect v0.0.0-20210303120416-fc811b0bcda2
	github.com/joho/godotenv v1.4.0
	github.com/streadway/amqp v1.0.0 // indirect
	golang.org/x/net v0.0.0-20220325170049-de3da57026de // indirect
	google.golang.org/grpc v1.32.0 // or newer
	google.golang.org/protobuf v1.28.0
)
