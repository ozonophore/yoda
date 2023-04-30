package event

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/yoda/common/pkg/eventbus"
	"github.com/yoda/webapp/internal/config"
	"github.com/yoda/webapp/internal/model/dto"
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
			logrus.Info("Receive mesasge: ", (string)(msg.Body))
			msg, _ := json.Marshal(&dto.RefreshAction{
				Type:    "refresh",
				NextRun: time.Now(),
			})
			Notify(msg)
		}
	})()

	RegistrationQueue("0", config.Consumer)
}

func CloseEvent() {
	eventbus.Close()
	logrus.Info("Event bus closed")
}
