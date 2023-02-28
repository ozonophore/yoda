package db

import (
	"github.com/yoda/app/pkg/client"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type DBProvider struct {
	client.EventListener
	db *gorm.DB
	tx *gorm.DB
}

func InitDatabase(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

func NewDbProvider(db *gorm.DB) *DBProvider {
	return &DBProvider{
		db: db,
	}
}

func (p *DBProvider) BeginOperation(t *string) *int64 {
	timeNow := time.Now()
	status := types.StatusTypeBegin
	m := model.Transaction{
		StartDate: &timeNow,
		Type:      t,
		Status:    &status,
	}
	p.db.Create(&m)
	p.tx = p.db.Begin()
	return m.ID
}

func (p *DBProvider) EndOperation(transaction *int64, status string) {
	switch status {
	case types.StatusTypeCompleted:
		d := p.tx.Commit()
		if d.Error != nil {
			log.Print("Commit was rejected %s", d.Error)
			status = types.StatusTypeRejected
		}
	default:
		p.tx.Rollback()
	}

	timeNow := time.Now()
	p.db.Updates(&model.Transaction{}).Where("id = ?", transaction).Updates(model.Transaction{
		EndDate: &timeNow,
		Status:  &status,
	})
}

func (p *DBProvider) WriteItem(model *model.StockItem) {
	p.tx.Create(&model)
}
