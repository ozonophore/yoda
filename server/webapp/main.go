package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/yoda/webapp/pkg/api"
	"github.com/yoda/webapp/pkg/config"
	"github.com/yoda/webapp/pkg/controller"
	"github.com/yoda/webapp/pkg/dao"
	"github.com/yoda/webapp/pkg/event"
	server2 "github.com/yoda/webapp/pkg/server"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx := context.Background()
	config, err := config.LoadConfig("config.yml")
	if err != nil {
		panic(err)
	}
	logger := createLogger(err, config)

	database := dao.InitDatabase(config.Database, logger)
	dao.NewRepositoryDAO(database)

	event.InitEvent(ctx, config.Mq)
	defer event.CloseEvent()

	server := controller.NewServerApi(&logger)

	router := api.HandlerWithOptions(server, api.GorillaServerOptions{
		BaseURL: "/api",
	})
	srv := server2.NewServer(config.Server, router)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error(err)
		}
	}()
	//tgbotContext, tgBotCancel := context.WithCancel(context.Background())
	//defer tgBotCancel()
	//go mqclient.StartTgBot(tgbotContext, config.TelegramBot, repository)
	logger.Infof("Server started on port %d", config.Server.Port)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	logger.Info("Shutting down")
	os.Exit(0)
}

func createLogger(err error, config *config.Config) log.Logger {
	logger := *log.StandardLogger()
	logger.Level, err = log.ParseLevel(config.LoggingLevel)
	log.SetLevel(logger.Level)
	if err != nil {
		logger.Printf("Error parsing log level: %s", err)
	}
	return logger
}
