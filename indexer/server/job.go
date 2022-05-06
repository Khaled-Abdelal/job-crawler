package server

import (
	"context"

	"github.com/Khaled-Abdelal/job-crawler/indexer/indexer"
	pb "github.com/Khaled-Abdelal/job-crawler/indexer/proto/jobservice"
	"github.com/elastic/go-elasticsearch/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedJobServiceServer
}

func (s *server) GetJobs(ctx context.Context, in *pb.GetJobsRequest) (*pb.GetJobsResponse, error) {
	esClient := ctx.Value(indexer.ElasticSearchClientContextKey).(*elasticsearch.Client)
	if esClient == nil {
		return nil, status.Error(codes.Internal, "No ElasticSearch connection found")
	}
	searchRes, _ := indexer.GetJobs(esClient, "front end", 0, 10)
	var jobs []*pb.Job
	for _, job := range searchRes.Jobs {
		jobs = append(jobs, &pb.Job{
			Title:       job.Title,
			URL:         job.URL,
			Source:      job.Source,
			Description: job.Description,
			Location:    job.Location,
			CompanyName: job.CompanyName,
		})
	}
	serverRes := pb.GetJobsResponse{
		Jobs:  jobs,
		Total: searchRes.Total,
	}
	return &serverRes, nil
}

func GetNewServer() *server {
	return &server{}
}
