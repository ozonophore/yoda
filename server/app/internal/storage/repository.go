package storage

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/configuration"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"strings"
	"time"
)

var repository *RepositoryDAO

type RepositoryDAO struct {
	db *gorm.DB
}

func stringToLogLevel(s *string) logger.LogLevel {
	switch strings.ToLower(*s) {
	case "info":
		return logger.Info
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	default:
		return logger.Silent
	}
}

func InitDatabase(config configuration.Database) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(stringToLogLevel(&config.LoggingLevel)),
	})
	if err != nil {
		logrus.Panic(err)
	}
	logrus.Info("Database log level: ", config.LoggingLevel)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	//err = db.Ping()
	if err != nil {
		log.Fatalln("Connection was rejected")
	}
	schema.RegisterSerializer("time", types.TimeSerializer{})
	return db
}

func NewRepositoryDAO(db *gorm.DB) *RepositoryDAO {
	if repository == nil {
		repository = &RepositoryDAO{
			db: db,
		}
	}
	return repository
}

func initIfError() {
	if repository == nil {
		logrus.Panic("Repository is not initialized")
	}
}

func CreateJob(job *model.Job) error {
	initIfError()
	tx := repository.db.WithContext(context.Background()).Create(job)
	return tx.Error
}

func GetJobs() (*[]model.Job, error) {
	initIfError()
	var jobs []model.Job
	err := repository.db.Find(&jobs).Error
	if err != nil {
		return nil, err
	}
	return &jobs, nil
}

func GetJob(id int) (*model.Job, error) {
	initIfError()
	var job model.Job
	err := repository.db.Preload("JobParameters").Preload("Owner").Where(`"id"=?`, id).First(&job).Error
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func GetJobWithOwnerByJobId(id int) (*model.Job, error) {
	initIfError()
	var job model.Job
	err := repository.db.Where(`"id"=?`, id).Find(&job).Error
	if err != nil {
		return nil, err
	}
	var params []model.OwnerMarketplace
	err = repository.db.Raw(`select om.* from "ml"."owner" o
			inner join "ml"."job_owner" jo on o."code" = jo."owner_code"
			inner join "ml"."owner_marketplace" om on o."code" = om."owner_code"
			where jo."job_id" = ? order by om."owner_code"`, id).Find(&params).Error
	if err != nil {
		return nil, err
	}
	job.Params = &params
	return &job, nil
}

func BeginOperation(jobId int) int64 {
	initIfError()
	timeNow := time.Now()
	status := types.StatusTypeBegin
	m := model.Transaction{
		StartDate: timeNow,
		Status:    status,
		JobId:     jobId,
	}
	repository.db.Create(&m)
	return m.ID
}

func EndOperation(transaction int64, status string) {
	initIfError()
	timeNow := time.Now()
	repository.db.Updates(&model.Transaction{}).Where("id = ?", transaction).Updates(model.Transaction{
		EndDate: &timeNow,
		Status:  status,
	})
}

func EndOperationWithMessage(transaction int64, status, message string) {
	initIfError()
	timeNow := time.Now()
	repository.db.Updates(&model.Transaction{}).Where("id = ?", transaction).Updates(model.Transaction{
		EndDate: &timeNow,
		Status:  status,
	})
}

func UpdatePrices(models *[]model.StockItem) error {
	initIfError()
	tx := repository.db.Begin()
	for _, model := range *models {
		err := tx.Exec(`UPDATE "dl"."stock" SET "barcode" = ?, "price" = ?, "discount" = ?, 
                   "price_after_discount" = ?, "card_created" = ?, "days_on_site" = ? WHERE "transaction_id" = ? AND "external_code" = ?`,
			model.Barcode, model.Price, model.Discount, model.PriceAfterDiscount, model.CardCreated, model.DaysOnSite, model.TransactionID, model.ExternalCode,
		).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func SelectUniqueStockItem(transactionId int64) *[]string {
	initIfError()
	var ex []string
	repository.db.Table(`"dl"."stock"`).Select(`"supplier_article"`).Distinct().Where(`"transaction_id" = ?`, transactionId).Find(&ex)
	return &ex
}

func SaveStocks(items *[]model.StockItem) error {
	initIfError()
	tx := repository.db.CreateInBatches(items, len(*items))
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func SaveStocksInBatches(items *[]*model.StockItem, batchSize int) error {
	initIfError()
	tx := repository.db.CreateInBatches(items, batchSize)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func SaveSales(items *[]model.Sale) error {
	initIfError()
	tx := repository.db.CreateInBatches(items, len(*items))
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func SaveOrders(items *[]model.Order) error {
	initIfError()
	tx := repository.db.CreateInBatches(items, len(*items))
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func SaveOrdersInBatches(items *[]model.Order, batchSize int) (int64, error) {
	initIfError()
	tx := repository.db.CreateInBatches(items, batchSize)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func UpdateAttributes(models *[]model.StockItem) error {
	initIfError()
	tx := repository.db.Begin()
	for _, model := range *models {
		tx.Model(&model).
			Select(`"subject"`, `"category"`, `"brand"`, `"name"`).
			Where(map[string]interface{}{`"transaction_id"`: model.TransactionID, `"supplier_article"`: *model.SupplierArticle}).
			UpdateColumns(model)
	}
	return tx.Commit().Error
}

func GetTlgEvents() (*[]model.TlgEvent, error) {
	initIfError()
	var events []model.TlgEvent
	err := repository.db.Find(&events).Error
	if err != nil {
		return nil, err
	}
	return &events, nil
}

func CreateOrUpdateLogLoad(model *model.LogLoad) error {
	initIfError()
	return repository.db.Save(model).Error
}

func GetLogLLoadByTrnsAndStatus(trns int64, status string) (*[]model.LogLoad, error) {
	initIfError()
	var logs []model.LogLoad
	err := repository.db.Where(`"transaction_id"=? and "status"=?`, trns, status).Find(&logs).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &logs, nil
		}
		return nil, err
	}
	return &logs, nil
}

func CallDailyData(transactionID int64) error {
	initIfError()
	return repository.db.Transaction(func(tx *gorm.DB) error {
		logrus.Debugf("call calc_order_delivered(%d)", transactionID)
		err := tx.Exec(`call calc_order_delivered(?)`, transactionID).Error
		if logrus.IsLevelEnabled(logrus.DebugLevel) && err == nil {
			logrus.Debug("call calc_order_delivered finished")
		}
		if err != nil {
			logrus.Errorf("call calc_order_delivered error: %s", err)
			return errors.Join(err)
		}
		logrus.Debugf("call calc_stock_daily(%d)", transactionID)
		err = tx.Exec(`call calc_stock_daily(?)`, transactionID).Error
		if logrus.IsLevelEnabled(logrus.DebugLevel) && err == nil {
			logrus.Debug("call calc_stock_daily finished")
		}
		if err != nil {
			logrus.Errorf("call calc_stock_daily error: %s", err)
			return errors.Join(err)
		}
		return nil
	})
}
