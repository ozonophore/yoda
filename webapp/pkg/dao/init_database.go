package dao

import (
	"github.com/sirupsen/logrus"
	"github.com/yoda/common/pkg/dbutil"
	"github.com/yoda/common/pkg/types"
	"github.com/yoda/webapp/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	loggerGorm "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDatabase(config config.Database, logger logrus.Logger) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{
		Logger: loggerGorm.Default.LogMode(dbutil.ParseLevel(config.LoggingLevel)),
	})
	if err != nil {
		logger.Panicln(err)
	}
	//sqlDB, _ := db.DB()
	//sqlDB.SetMaxIdleConns(10)
	//sqlDB.SetMaxOpenConns(10)
	//sqlDB.SetConnMaxLifetime(time.Hour)

	schema.RegisterSerializer("time", types.TimeSerializer{})
	return db
}
