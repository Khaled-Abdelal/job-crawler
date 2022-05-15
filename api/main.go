package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"

	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	pb "github.com/Khaled-Abdelal/job-crawler/indexer/proto/jobservice"
	"github.com/darahayes/go-boom"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	loadEnvFile()
	http.HandleFunc("/", serveFrontendHandler)

	http.HandleFunc("/api/jobs", searchJobsHandler)

	port := os.Getenv("SERVER_PORT")
	log.Printf("server listening at %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func searchJobsHandler(w http.ResponseWriter, r *http.Request) {
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
		conn, err := grpc.Dial(os.Getenv("INDEXER_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Error did not connect: %s", err)
			boom.Internal(w, "Error: internal server error")
			return
		}
		defer conn.Close()

		j := pb.NewJobServiceClient(conn)

		response, err := j.GetJobs(context.Background(), &pb.GetJobsRequest{SearchTerm: searchTerm, From: from, Size: size})
		if err != nil {
			log.Printf("Error when calling GetJobs: %s", err)
			boom.Internal(w, "Error: internal server error")
			return
		}
		r, _ := protojson.Marshal(response)
		w.Write([]byte(r))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}

func serveFrontendHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case "GET":
		proxy, err := NewProxy(os.Getenv("FRONTEND_URL"))
		if err != nil {
			log.Println("Error parsing frontend url", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		proxy.ServeHTTP(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}

func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(url), nil
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
