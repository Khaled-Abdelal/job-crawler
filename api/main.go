package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"

	pb "github.com/Khaled-Abdelal/job-crawler/indexer/proto/jobservice"
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
	switch r.Method {
	case "GET":
		searchTerm := r.URL.Query().Get("searchTerm")
		if searchTerm == "" {
			log.Print("search term missing")
			http.Error(w, "searchTerm is required", http.StatusBadRequest)
			return
		}

		var size int32 = 10
		var from int32 = 0
		if r.URL.Query().Get("size") != "" {
			strSize, err := strconv.Atoi(r.URL.Query().Get("size"))
			if err != nil {
				http.Error(w, "size must be valid int", http.StatusBadRequest)
				return
			}
			size = int32(strSize)
		}
		if r.URL.Query().Get("from") != "" {
			strFrom, err := strconv.Atoi(r.URL.Query().Get("from"))
			if err != nil {
				http.Error(w, "from must be valid int", http.StatusBadRequest)
				return
			}
			from = int32(strFrom)
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
		r, _ := json.Marshal(response)
		w.Write(r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}
