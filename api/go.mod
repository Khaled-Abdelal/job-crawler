module github.com/Khaled-Abdelal/job-crawler/api

go 1.16

replace github.com/Khaled-Abdelal/job-crawler/indexer/proto/jobservice => ./../indexer/proto/jobservice

require (
	github.com/Khaled-Abdelal/job-crawler/indexer v0.0.0-20220506163337-b19c81e14285
	google.golang.org/grpc v1.32.0
)
