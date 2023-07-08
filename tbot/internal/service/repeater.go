package service

import (
	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
	"github.com/yoda/tnot/internal/storage/notification"
	"time"
)

type repeaterStorageInterface interface {
	GetClients() (*[]int64, error)
	GetNotifications() (*[]notification.Notification, error)
	ConfirmNotification(id int64)
}

type repeaterNotifierInterface interface {
	GetSender() string
	SendReport(date time.Time, chatIds *[]int64) error
}

var (
	rep         repeaterStorageInterface
	logger      *logrus.Logger
	notificator repeaterNotifierInterface
)

func StartRepeater(repository repeaterStorageInterface, ntf repeaterNotifierInterface, log *logrus.Logger) {
	rep = repository
	notificator = ntf
	logger = log
	s := gocron.NewScheduler(time.UTC)
	s.Every(2).Minute().Do(repeat)
	s.StartAsync()
}

func repeat() {
	ntfc, err := rep.GetNotifications()
	if err != nil {
		logger.Errorf("Error while getting notifications: %v", err)
	}
	clients, err := rep.GetClients()
	if err != nil {
		logger.Error(err)
		return
	}
	for _, n := range *ntfc {
		sender := notificator.GetSender()
		if n.Sender != "all" && n.Sender != sender {
			return
		}
		switch {
		case n.Type == "report_yesterday":
			date, err := time.Parse(time.DateOnly, *n.Message)
			if err != nil {
				logger.Error(err)
				break
			}
			notificator.SendReport(date, clients)
			rep.ConfirmNotification(n.ID)
		}
	}
}
