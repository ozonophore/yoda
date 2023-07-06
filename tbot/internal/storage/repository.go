package storage

import (
	"github.com/sirupsen/logrus"
	"github.com/yoda/tnot/internal/configuration"
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

func (r *Repository) GetReport(date time.Time) (*[]report.Report, error) {
	var reports []report.Report
	err := r.db.Model(&report.Report{}).Where(`"report_date" = ?`, date).Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return &reports, nil
}

func (r *Repository) GetReportByCluster(date time.Time) (*[]report.ReportByCluster, error) {
	var reports []report.ReportByCluster
	err := r.db.Model(&report.ReportByCluster{}).Where(`"report_date" = ?`, date).Find(&reports).Error
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

func (r *Repository) GetChatIds() (*[]int64, error) {
	var values []int64
	err := r.db.Raw(`SELECT "chat_id" FROM ml."tg_group"`).Scan(&values).Error
	if err != nil {
		return nil, err
	}
	return &values, nil
}
