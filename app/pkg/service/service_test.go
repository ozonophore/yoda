package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/app/pkg/repository"
	"github.com/yoda/common/pkg/model"
	"gorm.io/gorm"
	"testing"
	"time"
)

func Test_Main(t *testing.T) {
	err := setup(t)
	assert.NoError(t, err, "setup error")

	config := configuration.InitConfig("config.yml")
	config.Database.Dsn = "postgres://root:secret@localhost:5432/pdb"
	config.Database.LoggingLevel = "error"
	config.Ozon.Host = "http://localhost:1080/ozon"
	config.Wb.Host = "http://localhost:1080/wb"
	database := repository.InitDatabase(config.Database)

	dao := repository.NewRepositoryDAO(database)
	assert.NotNil(t, dao, "repository is nil")
	jobId := 1
	transactionID := repository.BeginOperation(jobId)
	t.Run("WB", func(t *testing.T) {
		wbRun(t, config, transactionID, database)
	})
	t.Run("OZON", func(t *testing.T) {
		ozonRun(t, config, transactionID, database)
	})
}

func ozonRun(t *testing.T, config *configuration.Config, transactionID int64, database *gorm.DB) {
	ozonSservice := NewOzonService("OWNER", "111", "API_KEY", config)
	err := ozonSservice.Parsing(context.Background(), transactionID)
	assert.NoError(t, err, `Error after parsing: %s transaction %d`, err, transactionID)
	var count int64
	database.Table("stock").Where(`"transaction_id"=?`, transactionID).Count(&count)
	assert.Equal(t, int64(132), count, "count of stock")
	var stocks []model.StockItem
	database.Where(`"supplier_article"=? and "transaction_id"=?`, "ИР078795", transactionID).Find(&stocks)
	assert.Equal(t, 1, len(stocks), "count of stock")
	stock := stocks[0]
	assert.Equal(t, "ИР078795", *stock.SupplierArticle, "supplier_article")
	assert.Equal(t, "OZON", stock.Source, "source")
	assert.Equal(t, float64(1870), *stock.Price, "Price")
	assert.Equal(t, "5060244091740", *stock.Barcode, "barcode")
}

func wbRun(t *testing.T, config *configuration.Config, transactionID int64, database *gorm.DB) {
	wbService := NewWBService("OWNER", "token", config)
	err := wbService.Parsing(context.Background(), transactionID)
	assert.NoError(t, err, `Error after parsing: %s transaction %d`, err, transactionID)
	var count int64
	database.Table("stock").Where(`"transaction_id"=?`, transactionID).Count(&count)
	assert.Equal(t, int64(49), count, "count of stock")
	var stocks []model.StockItem
	database.Where(`"supplier_article"=? and "transaction_id"=?`, "ИР060045", transactionID).Find(&stocks)
	assert.Equal(t, 3, len(stocks), "count of stock")
	stock := stocks[0]
	assert.Equal(t, "ИР060045", *stock.SupplierArticle, "supplier_article")
	assert.Equal(t, "WB", stock.Source, "source")
	assert.Equal(t, float64(636), *stock.Price, "Price")
	lastChangeDate, _ := time.Parse(time.DateOnly, "2023-02-25")
	assert.Equal(t, lastChangeDate, stock.LastChangeDate, "LastChangeDate")
	lastChangeTime, _ := time.Parse(time.TimeOnly, "09:34:40")
	assert.Equal(t, lastChangeTime, stock.LastChangeTime, "LastChangeTime")
	assert.Equal(t, "4620003082726", *stock.Barcode, "barcode")
	database.Table("order").Where(`"transaction_id"=?`, transactionID).Count(&count)
	assert.Equal(t, int64(1391), count, "count of orders")
}

func setup(t *testing.T) error {
	dockerCompose, err := compose.NewDockerCompose("../../../docker/docker-compose.yml")
	if err != nil {
		t.Fatalf("compose.NewDockerCompose() error: %v", err)
	}

	t.Cleanup(func() {
		assert.NoError(t, dockerCompose.Down(context.Background(), compose.RemoveOrphans(true), compose.RemoveImagesLocal), "compose.Down()")
	})
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	dockerCompose.Up(ctx, compose.Wait(true))
	return err
}
