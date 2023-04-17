package event

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yoda/common/pkg/eventbus"
	"github.com/yoda/webapp/pkg/config"
)

var PUBLISHER_QUEUE_NAME string

func InitEvent(ctx context.Context, config config.Mq) {
	PUBLISHER_QUEUE_NAME = config.Publisher
	eventbus.NewBus(config.Url)
	eventbus.NewQueue(config.Consumer)
	eventbus.NewQueue(PUBLISHER_QUEUE_NAME)
	logrus.Info("Event bus initialized")

	//go eventbus.Consume(ctx, config.Consumer, func(msg amqp.Delivery) {
	//	switch msg.Type {
	//	case eventbus.EVENT_REGISTRATION:
	//		//var event eventbus.RegistrationRequest
	//		//json.Unmarshal(msg.Body, &event)
	//		//RegistrationQueue(&event, msg.MessageId)
	//	}
	//})()

	RegistrationQueue("0", config.Consumer)
}

func CloseEvent() {
	eventbus.Close()
	logrus.Info("Event bus closed")
}
