package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Khaled-Abdelal/job-crawler/indexer/indexer"
	pb "github.com/Khaled-Abdelal/job-crawler/indexer/proto/jobservice"
	"github.com/Khaled-Abdelal/job-crawler/indexer/server"
	"github.com/Khaled-Abdelal/job-crawler/indexer/worker"
	"github.com/Khaled-Abdelal/job-crawler/indexer/worker/consumers"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	loadEnvFile()
	esClient := indexer.NewElasticSearchClient()
	rabbitMQSession := worker.RabbitMQSetUp()                   // holds the rabbitMQ Connection, channel
	go consumers.CrawledJobsConsumer(rabbitMQSession, esClient) // listen for new crawled jobs to be indexed
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(ElasticSearchUnaryServerInterceptor(esClient)),
	)
	pb.RegisterJobServiceServer(s, server.GetNewServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// add Elasticsearch client to grpc server
func ElasticSearchUnaryServerInterceptor(client *elasticsearch.Client) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, indexer.ElasticSearchClientContextKey, client), req)
	}
}

func loadEnvFile() {
	env := os.Getenv("APP_ENV")
	if env == "production" {
		return
	}
	if env == "" {
		env = "development"
	}
	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
