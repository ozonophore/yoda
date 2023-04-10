package eventbus

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type ConsumeCallback func(msg amqp.Delivery)
type MessagesCallback func()

var Connection *amqp.Connection
var channel *amqp.Channel
var queues = make(map[string]*amqp.Queue)

func NewBus(url string) {
	var err error
	Connection, err = amqp.Dial(url)
	if err != nil {
		panic(err)
	}
	// opening a channel over the connection established to interact with RabbitMQ
	channel, err = Connection.Channel()
	if err != nil {
		panic(err)
	}
}

func NewQueue(queueName string) {
	_, ok := queues[queueName]
	if ok {
		return
	}
	queue, err := channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // args
	)
	if err != nil {
		logrus.Panic(err)
	}
	queues[queueName] = &queue
}

func Close() {
	channel.Close()
	Connection.Close()
}

func Consume(ctx context.Context, queueName string, callback ConsumeCallback) MessagesCallback {
	msgs, err := channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		panic(err)
	}
	return func() {
		for d := range msgs {
			select {
			case <-ctx.Done():
				logrus.Info("Closing consumer")
				return
			default:
				callback(d)
			}
		}
	}
}

func Publish(queueName string, body []byte) {
	err := channel.Publish(
		"",        // exchange
		queueName, // key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		panic(err)
	}
}
