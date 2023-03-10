package db

import (
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"strings"
	"time"
)

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

func InitDatabase(config *configuration.Configuration) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(stringToLogLevel(config.SqlLogger.Level)),
	})
	if err != nil {
		log.Fatalln(err)
	}
	//sqlDB, _ := db.DB()
	//sqlDB.SetMaxIdleConns(10)
	//sqlDB.SetMaxOpenConns(10)
	//sqlDB.SetConnMaxLifetime(time.Hour)

	//err = db.Ping()
	if err != nil {
		log.Fatalln("Connection was rejected")
	}
	return db
}

func NewRepositoryDAO(db *gorm.DB) *RepositoryDAO {
	return &RepositoryDAO{
		db: db,
	}
}

func (p *RepositoryDAO) BeginOperation(o, t *string, jobId *int) *int64 {
	timeNow := time.Now()
	status := types.StatusTypeBegin
	m := model.Transaction{
		StartDate: &timeNow,
		Source:    t,
		Status:    &status,
		OwnerCode: o,
		JobId:     jobId,
	}
	p.db.Create(&m)
	return m.ID
}

func (p *RepositoryDAO) Begin() *gorm.DB {
	return p.db.Begin()
}

func (p *RepositoryDAO) EndOperation(transaction *int64, status string) {
	timeNow := time.Now()
	p.db.Updates(&model.Transaction{}).Where("id = ?", transaction).Updates(model.Transaction{
		EndDate: &timeNow,
		Status:  &status,
	})
}

func (p *RepositoryDAO) UpdatePrices(models *[]model.StockItem) error {
	tx := p.db.Begin()
	for _, model := range *models {
		tx.Model(&model).
			Select("\"price\"", "\"discount\"", "\"price_after_discount\"").
			Where(map[string]interface{}{"\"transaction_id\"": *model.TransactionId, "\"external_code\"": *model.ExternalCode}).
			UpdateColumns(model)
	}
	return tx.Commit().Error
}

func (p *RepositoryDAO) SelectUniqueStockItem(transactionId *int64) *[]string {
	var ex []string
	p.db.Model(&model.StockItem{}).Distinct().Pluck("supplier_article", &ex).Where("transaction_id = ?", transactionId)
	return &ex
}

func (p *RepositoryDAO) SaveStocks(items *[]model.StockItem) error {
	tx := p.db.CreateInBatches(items, len(*items))
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (p *RepositoryDAO) UpdateAttributes(models *[]model.StockItem) error {
	tx := p.db.Begin()
	for _, model := range *models {
		tx.Model(&model).
			Select("\"subject\"", "\"category\"", "\"brand\"", "\"name\"").
			Where(map[string]interface{}{"\"transaction_id\"": *model.TransactionId, "\"supplier_article\"": *model.SupplierArticle}).
			UpdateColumns(model)
	}
	return tx.Commit().Error
}
