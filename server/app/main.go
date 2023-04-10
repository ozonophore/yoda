package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/app/pkg/event"
	"github.com/yoda/app/pkg/repository"
	"github.com/yoda/app/pkg/timer"
	"github.com/yoda/common/pkg/dao"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx := context.Background()
	config := configuration.InitConfig("config.yml")
	initLogging(config)
	database := repository.InitDatabase(config.Database)
	repository.NewRepositoryDAO(database)
	dao.NewDao(database)

	ctxCancel, cancel := context.WithCancel(ctx)
	event.InitEvent(ctxCancel, config.Mq)
	defer event.CloseEvent()
	//mq.NewConnection(config.Mq)
	//mq.NewConsumer(ctx, config.Mq, mqserver.HandlerReceive)

	scheduler := timer.NewScheduler(config)
	scheduler.InitJob()
	scheduler.Start()
	defer scheduler.Stop()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	waiting, cancel := context.WithTimeout(ctx, 5*time.Second)
	go func(ctx context.Context) {
		logrus.Infof("Waiting prosecces to finish")
		<-ctx.Done()
	}(waiting)
	defer cancel()
	logrus.Info("Shutting down")
}

func initLogging(config *configuration.Config) {
	lvl, err := logrus.ParseLevel(config.LoggingLevel)
	if err != nil {
		logrus.Panic(err)
	}
	logrus.SetLevel(lvl)
}
