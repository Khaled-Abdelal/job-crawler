# Indexer

A Golang service which handles indexing and searching jobs.

## Getting Started

### Using docker

- You need to have [docker](https://www.docker.com/) installed
  
- You need to have Both [RabbitMQ](https://hub.docker.com/_/rabbitmq) and [ElasticSearch](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html) running
  
- Build the image
  
    ```shell
   docker build -t indexer .
   ```

- Run the image

  ```shell
   docker run -p 50051:50051 indexer
   ```

## Libraries Used

- [go-elasticsearch](https://github.com/elastic/go-elasticsearch)
- [Go RabbitMQ Client Library](https://github.com/streadway/amqp)
- [GoDotEnv](https://github.com/joho/godotenv)
- [Google gRPC](https://pkg.go.dev/google.golang.org/grpc)
- [protobuf](https://pkg.go.dev/google.golang.org/protobuf)
