package stage

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/pipeline"
	"time"
)

type notifService interface {
	SetNotification(msg, sender, mtype string) error
}

type NotificationStep struct {
	service notifService
	logger  *logrus.Logger
	sender  string
}

func NewNotifyStep(service notifService, sender string, logg *logrus.Logger) *NotificationStep {
	return &NotificationStep{
		service: service,
		logger:  logg,
		sender:  sender,
	}
}

func (d *NotificationStep) Do(ctx context.Context, deps *map[string]pipeline.Stage, e error) (interface{}, error) {
	ps, ok := (*deps)["stock-daily-aggregator"]
	if !ok {
		return nil, errors.New("stock-daily-aggregator not found")
	}
	status := ps.GetStatus()
	date := status.Value.(time.Time)
	d.logger.Debugf("Notification: %s", date)
	return nil, d.service.SetNotification(date.Format(time.DateOnly), d.sender, "report_yesterday")
}
