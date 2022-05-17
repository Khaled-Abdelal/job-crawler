# Keywords

A Golang service which handles providing keywords to be crawled.

## Getting Started

### Using docker

- You need to have [docker](https://www.docker.com/) installed
  
- You need to have Both [RabbitMQ](https://hub.docker.com/_/rabbitmq) running.
  
- Build the image
  
    ```shell
   docker build -t keyword .
   ```

- Run the image

  ```shell
   docker run keyword
   ```

## Libraries Used

- [Go RabbitMQ Client Library](https://github.com/streadway/amqp)
- [gocron](https://github.com/go-co-op/gocron)
- [gorm](https://github.com/go-gorm/gorm)
- [postgres](https://github.com/go-gorm/postgres)
- [GoDotEnv](https://github.com/joho/godotenv)
