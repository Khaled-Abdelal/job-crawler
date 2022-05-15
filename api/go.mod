module github.com/Khaled-Abdelal/job-crawler/api

go 1.16

//replace github.com/Khaled-Abdelal/job-crawler/indexer/proto/jobservice => ./../indexer/proto/jobservice

require (
	github.com/Khaled-Abdelal/job-crawler/indexer v0.0.0-20220515052107-3e7abc0488e0 // indirect
	//github.com/Khaled-Abdelal/job-crawler/indexer v0.0.0-local
	github.com/darahayes/go-boom v0.0.0-20200826120415-fa5cb724143a
	github.com/joho/godotenv v1.4.0
	// github.com/Khaled-Abdelal/job-crawler/indexer v0.0.0-20220508053131-be2d74ee34f6
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	golang.org/x/sys v0.0.0-20220503163025-988cb79eb6c6 // indirect
	google.golang.org/genproto v0.0.0-20220505152158-f39f71e6c8f3 // indirect
	google.golang.org/grpc v1.46.0
	google.golang.org/protobuf v1.28.0
)

//replace github.com/Khaled-Abdelal/job-crawler/indexer v0.0.0-local => ../indexer

//replace github.com/Khaled-Abdelal/job-crawler/crawler v0.0.0-local => ../crawler
