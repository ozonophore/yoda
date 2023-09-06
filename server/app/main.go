package main

import (
	"context"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/go-co-op/gocron"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/configuration"
	"github.com/yoda/app/internal/event"
	"github.com/yoda/app/internal/integration"
	api "github.com/yoda/app/internal/integration/api"
	"github.com/yoda/app/internal/integration/dictionary"
	"github.com/yoda/app/internal/logging"
	"github.com/yoda/app/internal/pipeline"
	integration4 "github.com/yoda/app/internal/service/integration"
	service "github.com/yoda/app/internal/service/stock"
	"github.com/yoda/app/internal/stage"
	integration2 "github.com/yoda/app/internal/stage/integration"
	"github.com/yoda/app/internal/stage/product"
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
	repository := storage.NewRepository(database, config)

	ctxCancel, cancel := context.WithCancel(ctx)
	event.InitEvent(ctxCancel, config.Mq)
	defer event.CloseEvent()

	uo := integration.NewUpdaterOrganisations(config.Integration)
	dictionary.InitDictionary(config.Integration)

	scheduler := timer.NewScheduler(config, repository)
	scheduler.Subscribe(event.CreateObserver())
	event.Subscribe(scheduler)
	event.SubscribeOrg(uo)
	scheduler.InitJob()
	startStockAggregator(database, scheduler.GetScheduler(), config)
	client, _ := InitClient(config)
	StartInitializer(repository, scheduler.GetScheduler(), client)

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

func InitClient(config *configuration.Config) (api.ClientWithResponsesInterface, error) {
	apiKeyProvider, _ := securityprovider.NewSecurityProviderApiKey("header", "Key", config.Integration.Token)
	return api.NewClientWithResponses(config.Integration.Host,
		logging.WithLoggerIntegrationFn(config.Integration.LogLevel),
		api.WithRequestEditorFn(apiKeyProvider.Intercept),
	)
}

func StartInitializer(repository *storage.Repository, sch *gocron.Scheduler, client api.ClientWithResponsesInterface) *stage.Initializer {
	stage.Register(3, func() pipeline.Stage {
		service := integration4.NewStockService(repository)
		return integration2.NewStockStage(service, client, logrus.StandardLogger())
	})

	factory := stage.NewStageFactory()
	initializer := stage.NewInitializer(repository, factory, sch)
	initializer.Do(3)

	sch.Every(1).Minute().Do(func() {
		initializer.Repeat()
	})
	return initializer
}

func startStockAggregator(db *gorm.DB, sch *gocron.Scheduler, config *configuration.Config) {
	rep := storage.NewRepository(db, config)
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
	logger := logrus.StandardLogger()
	srv := service.NewStockService(rep, logger)
	interceptor := logging.NewInterceptor(logrus.StandardLogger())
	step := stock.NewDailyStep(srv, logger)
	stg := pipeline.NewSimpleStageWithTag(step, "stock-daily-aggregator").AddSubscriber(interceptor)
	defStage := pipeline.NewSimpleStageWithTag(stock.NewDefectureStep(srv, logger), "stock-defecture-aggregator").AddSubscriber(interceptor)
	stg.AddNext(defStage)
	repStage := pipeline.NewSimpleStageWithTag(stock.NewReportStep(srv, logger), "stock-report-aggregator").AddSubscriber(interceptor)
	defStage.AddNext(repStage)

	stageDefCluster := pipeline.NewSimpleStageWithTag(stock.NewDefByClustersStep(srv, logger), "def-cluster").AddSubscriber(interceptor)
	stageRepCluster := pipeline.NewSimpleStageWithTag(stock.NewReportByClustersStep(srv, logger), "rep-cluster").AddSubscriber(interceptor)
	stg.AddNext(stageDefCluster.AddNext(stageRepCluster))

	//Отчет по позициям
	stageDefItem := pipeline.NewSimpleStageWithTag(stock.NewDefByItemStep(srv, logger), "def-item").AddSubscriber(interceptor)
	stageRepItem := pipeline.NewSimpleStageWithTag(stock.NewReportByItemStep(srv, logger), "rep-item").AddSubscriber(interceptor)
	stg.AddNext(stageDefItem.AddNext(stageRepItem))

	stageNotification := pipeline.NewSimpleStageWithTag(stage.NewNotifyStep(srv, config.Sender, logger), "notification").AddSubscriber(interceptor)
	stageRepCluster.AddNext(stageNotification)
	repStage.AddNext(stageNotification)
	stageRepItem.AddNext(stageNotification)

	stageProduct := pipeline.NewSimpleStageWithTag(product.NewProductStep(srv, logger), "report-by-product").AddSubscriber(interceptor)
	stg.AddNext(stageProduct)
	stageProduct.AddNext(stageNotification)

	sj, err := sch.Every(1).Day().At(atTime).Tag("2").Do(func() {
		pipe := pipeline.NewPipeline()
		err := pipe.Do(context.Background(), stg).Error()
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
