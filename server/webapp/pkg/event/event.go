package event

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/yoda/common/pkg/eventbus"
	"github.com/yoda/webapp/pkg/config"
	"time"
)

var PUBLISHER_QUEUE_NAME string
var CONSUMER_QUEUE_NAME string

var responseQueue = make(map[string]chan string)

func InitEvent(ctx context.Context, config config.Mq) {
	PUBLISHER_QUEUE_NAME = config.Publisher
	CONSUMER_QUEUE_NAME = config.Consumer
	eventbus.NewBus(config.Url)
	eventbus.NewQueue(config.Consumer)
	eventbus.NewQueue(PUBLISHER_QUEUE_NAME)
	logrus.Info("Event bus initialized")

	go eventbus.Consume(ctx, config.Consumer, func(msg amqp.Delivery) {
		switch msg.Type {
		case eventbus.EVENT_REGISTRATION:
			//var event eventbus.RegistrationRequest
			//json.Unmarshal(msg.Body, &event)
			//RegistrationQueue(&event, msg.MessageId)
		case eventbus.EVENT_RUN_JOB:
			var response eventbus.MessageRunTaskResponse
			json.Unmarshal(msg.Body, &response)
			select {
			case responseQueue[response.ID] <- response.ID:
				logrus.Info("Run job success for id: ", response.ID)
			case <-time.After(30 * time.Second):
				logrus.Error("Run job timeout")
			}
		}
	})()

	RegistrationQueue("0", config.Consumer)
}

func RegistrationHandlerRunJob(id string) {
	c := make(chan string)
	responseQueue[id] = c
	select {
	case <-c:
		logrus.Info("Run job success for id: ", id)
	case <-time.After(30 * time.Second):
		logrus.Error("Message was not received after 30 seconds")
	}
}

func CloseEvent() {
	eventbus.Close()
	logrus.Info("Event bus closed")
}
