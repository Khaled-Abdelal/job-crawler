# Crawler

A Golang service which handles receiving keywords and crawling different job sites for jobs then pushing them to be indexed.

## Getting Started

### Using docker

- You need to have [docker](https://www.docker.com/) installed.
  
- You need to have Both [RabbitMQ](https://hub.docker.com/_/rabbitmq) running.
  
- Build the image.
  
    ```shell
   docker build -t crawler .
   ```

- Run the image.

  ```shell
   docker run crawler
   ```

## Libraries Used

- [Colly](https://github.com/gocolly/colly)
- [Go RabbitMQ Client Library](https://github.com/streadway/amqp)
- [GoDotEnv](https://github.com/joho/godotenv)
