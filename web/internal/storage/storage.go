package storage

import (
	"github.com/sirupsen/logrus"
	"github.com/yoda/web/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"strings"
	"time"
)

type Storage struct {
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

func NewStorage(config config.Database) *Storage {
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(stringToLogLevel(&config.LoggingLevel)),
	})
	if err != nil {
		logrus.Panic(err)
	}
	logrus.Info("Database log level: ", config.LoggingLevel)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	//err = db.Ping()
	if err != nil {
		log.Fatalln("Connection was rejected")
	}
	return &Storage{
		db: db,
	}
}
