package worker

import (
	"context"
	"os"

	"github.com/isayme/go-amqp-reconnect/rabbitmq"
)

// create a unique identifier for the context key
type key int

const ampqSessionKeyID key = iota

// ampqSession composes an rabbitmq.Connection with an rabbitmq.Channel
type ampqSession struct {
	*rabbitmq.Connection
	*rabbitmq.Channel
}

func RabbitMQSetUp(ctx context.Context) context.Context {
	connection := connect(os.Getenv("RABBIT_MQ_CONNECTION_STRING"))
	channelRabbitMQ, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	ctx = context.WithValue(ctx, ampqSessionKeyID, ampqSession{Connection: connection, Channel: channelRabbitMQ})
	return ctx
}

func GetSessionFromContext(ctx context.Context) ampqSession {
	return ctx.Value(ampqSessionKeyID).(ampqSession)
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
