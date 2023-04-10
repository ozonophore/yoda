package event

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/common/pkg/eventbus"
)

func InitEvent(ctx context.Context, config configuration.Mq) {
	eventbus.NewBus(config.Url)
	eventbus.NewQueue(config.Consumer)
	logrus.Info("Event bus initialized")
	go eventbus.Consume(ctx, config.Consumer, func(msg amqp.Delivery) {

	})()
}

func CloseEvent() {
	eventbus.Close()
	logrus.Info("Event bus closed")
}
