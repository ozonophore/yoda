package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yoda/tnot/internal/configuration"
	"github.com/yoda/tnot/internal/service"
	"github.com/yoda/tnot/internal/storage"
	"os"
	"os/signal"
	"time"
)

func main() {
	println(time.Now().AddDate(0, 0, -1).Format("2006-01-02"))
	os.Exit(0)
	config := configuration.InitConfig("config.yml")
	repository := storage.NewRepository(config.Database)
	reportService := service.NewReportService(repository)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go service.StartBot(config.Token, config.WorkDir, ctx, reportService, repository)

	// Tell the user the bot is online
	logrus.Println("Start listening for updates. Press enter to stop")

	// Wait for a newline symbol, then cancel handling updates
	//bufio.NewReader(os.Stdin).ReadBytes('\n')
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	cancel()
	logrus.Info("Shutting down")
}
