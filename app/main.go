package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/app/pkg/mqserver"
	"github.com/yoda/app/pkg/repository"
	"github.com/yoda/app/pkg/timer"
	"github.com/yoda/common/pkg/mq"
	"log"
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

	mq.NewConnection(config.Mq)
	mq.NewConsumer(ctx, config.Mq, mqserver.HandlerReceive)

	scheduler := timer.InitScheduler(&ctx, config)
	scheduler.StartAsync()
	logrus.Info("Scheduler started")
	defer func() {
		scheduler.Stop()
		logrus.Info("Scheduler stopped")
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Println("Shutting down")
	os.Exit(0)
}

func initLogging(config *configuration.Config) {
	lvl, err := logrus.ParseLevel(config.LoggingLevel)
	if err != nil {
		logrus.Panic(err)
	}
	logrus.SetLevel(lvl)
}
