package main

import (
	"context"
	"github.com/go-co-op/gocron"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/configuration"
	"github.com/yoda/app/internal/event"
	"github.com/yoda/app/internal/integration"
	"github.com/yoda/app/internal/integration/dictionary"
	"github.com/yoda/app/internal/logging"
	"github.com/yoda/app/internal/pipeline"
	service "github.com/yoda/app/internal/service/stock"
	stage2 "github.com/yoda/app/internal/stage"
	"github.com/yoda/app/internal/stage/stock"
	"github.com/yoda/app/internal/storage"
	"github.com/yoda/app/internal/timer"
	"github.com/yoda/common/pkg/dao"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"sort"
	"strings"
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
	startStockAggregator(database, scheduler.GetScheduler(), config)

	scheduler.Start()
	defer scheduler.StopAll()
	scheduler.RunImmediately(2)

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

func startStockAggregator(db *gorm.DB, sch *gocron.Scheduler, config *configuration.Config) {
	rep := storage.NewRepository(db)
	j, err := rep.GetJob(2)
	if !j.IsActive {
		logrus.Infof("Job with tag(%d) is not active", 2)
		return
	}
	if err != nil {
		logrus.Panicf("Error while getting job(%d): %v", 2, err)
	}
	job := timer.JobByTag(sch, 2)
	atTime := strings.ReplaceAll(*j.AtTime, ",", ";")
	if job != nil {
		logrus.Infof("Job with tag(%d) already exists", 2)
		actualAtTimes := job.ScheduledAtTimes()
		expectedAtTimes := strings.Split(atTime, ";")
		sort.Strings(actualAtTimes)
		sort.Strings(expectedAtTimes)
		if strings.Join(actualAtTimes, ";") == strings.Join(expectedAtTimes, ";") {
			return
		}
		err := sch.RemoveByTag("2")
		if err != nil {
			logrus.Errorf("Error while removing job with tag(%d): %v", 2, err)
		}
	}
	srv := service.NewStockService(rep)
	interceptor := logging.NewInterceptor(logrus.StandardLogger())
	logger := logrus.StandardLogger()
	step := stock.NewDailyStep(srv, logger)
	stage := pipeline.NewSimpleStageWithTag(step, "stock-daily-aggregator").AddSubscriber(interceptor)
	defStage := pipeline.NewSimpleStageWithTag(stock.NewDefectureStep(srv, logger), "stock-defecture-aggregator").AddSubscriber(interceptor)
	stage.AddNext(defStage)
	repStage := pipeline.NewSimpleStageWithTag(stock.NewReportStep(srv, logger), "stock-report-aggregator").AddSubscriber(interceptor)
	defStage.AddNext(repStage)

	stageDefCluster := pipeline.NewSimpleStageWithTag(stock.NewDefByClustersStep(srv, logger), "def-cluster").AddSubscriber(interceptor)
	stageRepCluster := pipeline.NewSimpleStageWithTag(stock.NewReportByClustersStep(srv, logger), "rep-cluster").AddSubscriber(interceptor)
	stage.AddNext(stageDefCluster.AddNext(stageRepCluster))

	stageNotification := pipeline.NewSimpleStageWithTag(stage2.NewNotifyStep(srv, config.Sender, logger), "notification").AddSubscriber(interceptor)
	stageRepCluster.AddNext(stageNotification)
	repStage.AddNext(stageNotification)

	sj, err := sch.Every(1).Day().At(atTime).Tag("2").Do(func() {
		pipe := pipeline.NewPipeline()
		err := pipe.Do(context.Background(), stage).Error()
		if err != nil {
			logrus.Errorf("Error while running stock aggregator: %v", err)
		}
	})
	if err != nil {
		logrus.Panicf("Error while scheduling job: %v", err)
	}
	sj.SetEventListeners(func() {
		logrus.Infof("Job with tag(%d) is running", 2)
	}, func() {
		logrus.Infof("Job with tag(%d) is done. Next run: %s", 2, sj.NextRun().UTC())
	})
}
