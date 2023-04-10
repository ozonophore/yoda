package mq

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel
var qRead amqp.Queue
var qWrite amqp.Queue

type Mq struct {
	Url        string `koanf:"url"`
	ReadQueue  string `koanf:"read_queue"`
	WriteQueue string `koanf:"write_queue"`
	MaxLength  int32  `koanf:"max_length"`
}

func NewConnection(config Mq) error {
	var err error
	conn, err = amqp.Dial(config.Url)
	if err != nil {
		return err
	}
	ch, err = conn.Channel()
	if err != nil {
		return err
	}

	args := make(amqp.Table)
	args["x-max-length"] = config.MaxLength

	// Create a Queue to send the message to.
	qWrite, err = ch.QueueDeclare(
		config.WriteQueue, // name
		false,             // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		args,              // arguments
	)
	if err != nil {
		return err
	}
	return nil
}

func NewConsumer(ctx context.Context, config Mq, handleMessage func(msg amqp.Delivery)) error {

	args := make(amqp.Table)
	args["x-max-length"] = config.MaxLength
	// Create a Queue to receive the message to.
	qRead, err := ch.QueueDeclare(
		config.ReadQueue, // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		args,             // arguments
	)
	if err != nil {
		return errors.Join(err)
	}

	msgs, err := ch.Consume(
		qRead.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return err
	}
	go func() {
		for msg := range msgs {
			select {
			case <-ctx.Done():
				logrus.Info("Consumer bot stopped")
			default:
				handleMessage(msg)
			}
		}
	}()
	logrus.Info("Consumer started")
	return nil
}

func SendMessageText(msg string) error {
	err := ch.Publish(
		"",          // exchange
		qWrite.Name, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	logrus.Debugf(" [x] Sent %s", msg)
	if err != nil {
		logrus.Error(err)
	}
	return err
}

func SendMessage(dataType string, data []byte) error {
	headers := amqp.Table{
		"message_type": dataType,
	}
	err := ch.Publish(
		"",          // exchange
		qWrite.Name, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "json/application",
			Body:        data,
			Headers:     headers,
		})
	logrus.Debugf(" [x] Sent %s", string(data))
	if err != nil {
		logrus.Error(err)
	}
	return err
}

func Close() {
	ch.Close()
	conn.Close()
}
