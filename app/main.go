package main

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/app/pkg/event"
	jobf "github.com/yoda/app/pkg/job"
	"github.com/yoda/app/pkg/mqserver"
	"github.com/yoda/app/pkg/repository"
	"github.com/yoda/common/pkg/mq"
	"github.com/yoda/common/pkg/types"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"
)

func main() {
	ctx := context.Background()
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.SingletonModeAll()

	config := configuration.InitConfig("config.yml")
	initLogging(config)
	database := repository.InitDatabase(config.Database)
	repository.NewRepositoryDAO(database)
	mq.NewConnection(config.Mq)
	mq.NewConsumer(ctx, config.Mq, mqserver.HandlerReceive)

	_, err := scheduler.Every(1).Minute().Do(runETL, config, ctx)
	if err != nil {
		log.Panicf("Error after create job: %s", err)
	}
	scheduler.StartAsync()

	defer scheduler.Stop()

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

func runETL(config *configuration.Config, ctx context.Context) {
	jobs, err := repository.GetJobs()
	if err != nil {
		log.Panicf("Error after get jobs: %s", err)
	}
	var info []string
	for _, job := range *jobs {
		owner := &job.OwnerCode
		transactionID := repository.BeginOperation(*owner, job.ID)
		var err error
		logrus.Info("Start parsing for owner: ", *owner)
		var mrks []string
		for _, param := range job.JobParameters {
			var loader jobf.DataLoader
			loader, err = jobf.JobFactory(param.Source, *owner, *param.Password, param.ClientID, config)
			if err != nil {
				logrus.Errorf("Error after lookup a loader: %s", err)
				continue
			}
			newContext, _ := context.WithTimeout(ctx, time.Duration(config.Timeout)*time.Second)
			err = loader.Parsing(newContext, transactionID)
			if err != nil {
				logrus.Errorf("Error after parsing: %s", err)
				continue
			}
			mrks = append(mrks, fmt.Sprintf("%s", param.Source))
		}
		if err == nil {
			repository.EndOperation(transactionID, types.StatusTypeCompleted)
		} else {
			repository.EndOperation(transactionID, types.StatusTypeRejected)
		}
		logrus.Info("Finish parsing for owner: ", *owner)
		info = append(info, fmt.Sprintf("Owner %s marketplaces: %s", job.Owner.Name, strings.Join(mrks, ",")))
	}
	go event.ProcessInfo(&info)
}
