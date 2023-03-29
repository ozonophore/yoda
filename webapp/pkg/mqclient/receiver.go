package mqclient

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/yoda/common/pkg/mq"
	"github.com/yoda/webapp/pkg/dao"
	"strings"
)

func HandleMessage(msg amqp.Delivery) {
	msgType := msg.Headers["message_type"]
	switch msgType {
	case mq.HEADER_ETL_INFO:
		data := mq.MessageETLInfoResponse{}
		if err := json.Unmarshal(msg.Body, &data); err != nil {
			logrus.Error(err)
		}
		SendETLInfo(data)
	}
}

func SendETLInfo(data mq.MessageETLInfoResponse) {
	msg, _ := dao.GetTlgMessageById(data.ID)
	if msg == nil {
		return
	}
	SendMessage(msg.ChatID, msg.MessageID, fmt.Sprintf(`%s`, strings.Join(data.Data, ",\n")))
}
