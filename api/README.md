# API Gateway

A Golang front facing server which handles all incoming traffic.

## Getting Started

### Using docker

- You need to have [docker](https://www.docker.com/) installed
  
- Build the image
  
    ```shell
   docker build -t api .
   ```

- Run the image

  ```shell
   docker run -p 8081:8081 api
   ```

## Libraries Used

- [Google gRPC](https://pkg.go.dev/google.golang.org/grpc)
- [protobuf](https://pkg.go.dev/google.golang.org/protobuf)
- [Boom](https://github.com/darahayes/go-boom)
- [GoDotEnv](https://github.com/joho/godotenv)
