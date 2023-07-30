package storage

import (
	"github.com/sirupsen/logrus"
	"github.com/yoda/tnot/internal/configuration"
	"github.com/yoda/tnot/internal/storage/notification"
	"github.com/yoda/tnot/internal/storage/report"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"strings"
	"time"
)

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

type Repository struct {
	db *gorm.DB
}

func NewRepository(config configuration.Database) *Repository {
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(stringToLogLevel(&config.LoggingLevel)),
	})
	if err != nil {
		panic(err)
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
	schema.RegisterSerializer("time", TimeSerializer{})
	return &Repository{db: db}
}

func (r *Repository) GetReport(date time.Time) (*[]map[string]interface{}, error) {
	//var reports []report.Report
	//err := r.db.Model(&report.Report{}).Where(`"report_date" = ?`, date.Format(time.DateOnly)).Find(&reports).Error

	var results []map[string]interface{}
	err := r.db.Table("dl.sales_stock").Where(`"report_date" = ?`, date.Format(time.DateOnly)).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (r *Repository) GetReportByCluster(date time.Time) (*[]report.ReportByCluster, error) {
	var reports []report.ReportByCluster
	err := r.db.Model(&report.ReportByCluster{}).Where(`"report_date" = ?`, date.Format(time.DateOnly)).Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return &reports, nil
}

func (r *Repository) GetReportByItem(date time.Time) (*[]report.ReportByItem, error) {
	var reports []report.ReportByItem
	err := r.db.Model(&report.ReportByItem{}).Where(`"report_date" = ?`, date.Format(time.DateOnly)).Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return &reports, nil
}

func (r *Repository) AddGroup(userName, groupName string, chatId int64) error {
	return r.db.Exec(`INSERT INTO ml."tg_group" ("user_name", "group_name", "chat_id") VALUES (?, ?, ?)`, userName, groupName, chatId).Error
}

func (r *Repository) DeleteGroup(userName, groupName string) error {
	return r.db.Exec(`DELETE FROM ml."tg_group" WHERE "user_name" = ? AND "group_name" = ?`, userName, groupName).Error
}

func (r *Repository) GetClients() (*[]int64, error) {
	var values []int64
	err := r.db.Raw(`SELECT "chat_id" FROM ml."tg_group"`).Scan(&values).Error
	if err != nil {
		return nil, err
	}
	return &values, nil
}

func (r *Repository) GetNotifications() (*[]notification.Notification, error) {
	var values []notification.Notification
	err := r.db.Model(&notification.Notification{}).Where(`"is_sent" = false`).Scan(&values).Error
	if err != nil {
		return nil, err
	}
	return &values, nil
}

func (r *Repository) ConfirmNotification(id int64) {
	err := r.db.Exec(`UPDATE ml."notification" SET "is_sent" = true WHERE "id" = ?`, id)
	if err.Error != nil {
		logrus.Error(err.Error)
	}
}
