package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/yoda/webapp/internal/config"
	"github.com/yoda/webapp/internal/dao"
	"github.com/yoda/webapp/internal/event"
	"github.com/yoda/webapp/internal/server"
	"github.com/yoda/webapp/internal/ws"
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

	ws := ws.StartServer()
	event.InitEvent(ctx, config.Mq)
	event.AddObserver(ws)
	defer event.CloseEvent()

	s := server.StartServer(config.Server, &logger)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.Shutdown(ctx)
	ws.Close(ctx)
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
