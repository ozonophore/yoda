package event

import (
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/pkg/mqserver"
	"github.com/yoda/app/pkg/repository"
	"github.com/yoda/common/pkg/mq"
)

func ProcessInfo(data *[]string) {
	events, err := repository.GetTlgEvents()
	if err != nil {
		logrus.Error(err)
		return
	}
	for _, event := range *events {
		if err := mqserver.SendMessageEtlInfoResponse(mq.MessageETLInfoResponse{
			ID:   event.ChatID,
			Data: *data,
		}); err != nil {
			logrus.Error(err)
		}
	}
}
