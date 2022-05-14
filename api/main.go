package main

import (
	"context"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"

	"google.golang.org/protobuf/encoding/protojson"

	pb "github.com/Khaled-Abdelal/job-crawler/indexer/proto/jobservice"
	"github.com/darahayes/go-boom"
	"google.golang.org/grpc"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/api/jobs", jobsHandler)

	log.Printf("server listening at %d", 8081)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func jobsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case "GET":
		searchTerm := r.URL.Query().Get("searchTerm")
		if searchTerm == "" {
			boom.BadRequest(w, "searchTerm is required")
			return
		}

		var size int32 = 10
		var from int32 = 0
		if r.URL.Query().Get("size") != "" {
			strSize, err := strconv.Atoi(r.URL.Query().Get("size"))
			if err != nil {
				boom.BadRequest(w, "size must be valid int")
				return
			}
			size = int32(strSize)
			if size < 0 {
				boom.BadRequest(w, "size must be positive")
				return
			}
		}
		if r.URL.Query().Get("from") != "" {
			strFrom, err := strconv.Atoi(r.URL.Query().Get("from"))
			if err != nil {
				boom.BadRequest(w, "from must be valid int")
				return
			}
			from = int32(strFrom)
			if from < 0 {
				boom.BadRequest(w, "from must be positive")
				return
			}
		}
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(":50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		j := pb.NewJobServiceClient(conn)

		response, err := j.GetJobs(context.Background(), &pb.GetJobsRequest{SearchTerm: searchTerm, From: from, Size: size})
		if err != nil {
			log.Fatalf("Error when calling GetJobs: %s", err)
		}
		r, _ := protojson.Marshal(response)
		w.Write([]byte(r))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}
