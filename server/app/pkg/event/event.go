package event

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/common/pkg/eventbus"
)

var queues []string

func InitEvent(ctx context.Context, config configuration.Mq) {
	eventbus.NewBus(config.Url)
	eventbus.NewQueue(config.Consumer)
	logrus.Info("Event bus initialized")

	go eventbus.Consume(ctx, config.Consumer, func(msg amqp.Delivery) {
		switch msg.Type {
		case eventbus.EVENT_REGISTRATION:
			var event eventbus.RegistrationRequest
			json.Unmarshal(msg.Body, &event)
			RegistrationQueue(&event, msg.MessageId)
		}
	})()
}

func PublishToAll(body *[]byte, msgType, msgId string) {
	for _, queue := range queues {
		eventbus.Publish(queue, msgType, msgId, *body)
	}
}

func AddQueue(name string) {
	eventbus.NewQueue(name)
	queues = append(queues, name)
}

func CloseEvent() {
	eventbus.Close()
	logrus.Info("Event bus closed")
}
