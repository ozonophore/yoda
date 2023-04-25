package event

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/app/pkg/listener"
	"github.com/yoda/common/pkg/eventbus"
	"gorm.io/gorm/utils"
	"time"
)

var queues []string
var lstrn *listener.Listener

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
		case eventbus.EVENT_RUN_JOB:
			var req eventbus.MessageRunTaskRequest
			json.Unmarshal(msg.Body, &req)
			if time.Now().Sub(req.Date).Seconds() > 30 {
				logrus.Error("Run job timeout")
				return
			}
			AddQueue(req.QueueName)
			logrus.Info("Run job success for id: ", req.ID)
			if lstrn != nil {
				(*lstrn).RunTask(req.JobId)
			}
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
	if utils.Contains(queues, name) {
		return
	}
	queues = append(queues, name)
}

func SetListener(listener *listener.Listener) {
	lstrn = listener
}

func CloseEvent() {
	eventbus.Close()
	logrus.Info("Event bus closed")
}
