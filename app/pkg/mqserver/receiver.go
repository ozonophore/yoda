package mqserver

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/yoda/common/pkg/mq"
)

func HandlerReceive(msg amqp.Delivery) {
	msgType := msg.Headers["message_type"]
	switch msgType {
	case mq.HEADER_ETL_INFO:
		data := mq.MessageETLInfoRequest{}
		logrus.Debugf(`[x] Reseived %s`, string(msg.Body))
		if err := json.Unmarshal(msg.Body, &data); err != nil {
			logrus.Error(err)
		}
		handlerETLInfo(data)
	}
}

func handlerETLInfo(data mq.MessageETLInfoRequest) {
	// Nothing to do
}
