package worker

import (
	"os"

	"github.com/isayme/go-amqp-reconnect/rabbitmq"
)

// ampqSession composes an rabbitmq.Connection with an rabbitmq.Channel
type AMPQSession struct {
	*rabbitmq.Connection
	*rabbitmq.Channel
}

func RabbitMQSetUp() *AMPQSession {
	connection := connect(os.Getenv("RABBIT_MQ_CONNECTION_STRING"))
	channelRabbitMQ, err := connection.Channel()
	if err != nil {
		panic("Can't connect to RabbitMQ")
	}
	return &AMPQSession{Connection: connection, Channel: channelRabbitMQ}
}

func connect(url string) *rabbitmq.Connection {
	connectRabbitMQ, err := rabbitmq.Dial(url)
	if err != nil {
		panic(err)
	}
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	createQueues(channelRabbitMQ)
	return connectRabbitMQ
}

func createQueues(channel *rabbitmq.Channel) error {
	defer channel.Close()
	_, err := channel.QueueDeclare(
		os.Getenv("CRAWLED_JOBS_QUEUE"), // queue name
		true,                            // durable
		false,                           // auto delete
		false,                           // exclusive
		false,                           // no wait
		nil,                             // arguments
	)
	if err != nil {
		panic(err)
	}
	return nil
}
