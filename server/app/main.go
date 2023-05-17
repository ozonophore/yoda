package main

import (
	"context"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/configuration"
	"github.com/yoda/app/internal/event"
	"github.com/yoda/app/internal/integration"
	"github.com/yoda/app/internal/integration/dictionary"
	"github.com/yoda/app/internal/storage"
	"github.com/yoda/app/internal/timer"
	"github.com/yoda/common/pkg/dao"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx := context.Background()
	config := configuration.InitConfig("config.yml")
	initLogging(config)
	database := storage.InitDatabase(config.Database)
	storage.NewRepositoryDAO(database)
	dao.NewDao(database)

	ctxCancel, cancel := context.WithCancel(ctx)
	event.InitEvent(ctxCancel, config.Mq)
	defer event.CloseEvent()

	uo := integration.NewUpdaterOrganisations(config.Integration)
	dictionary.InitDictionary(config.Integration)

	scheduler := timer.NewScheduler(config)
	scheduler.Subscribe(event.CreateObserver())
	event.Subscribe(scheduler)
	event.SubscribeOrg(uo)
	scheduler.InitJob()
	scheduler.Start()
	defer scheduler.StopAll()

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
	pathMap := "./log/log.log"
	logrus.AddHook(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))
	logrus.SetLevel(lvl)
}
