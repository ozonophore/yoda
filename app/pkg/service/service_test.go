package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/app/pkg/repository"
	"github.com/yoda/common/pkg/model"
	"testing"
	"time"
)

func Test_Main(t *testing.T) {
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
	//assert.NoError(t, dockerCompose.Up(ctx, compose.Wait(true)), "compose.Up()")

	config := configuration.InitConfig()
	//config.SqlLogger.Level = "info"
	config.Dsn = "postgres://root:secret@localhost:5432/pdb"
	database := repository.InitDatabase(config)

	repository := repository.NewRepositoryDAO(database)
	assert.NotNil(t, repository, "repository is nil")
	wbService := NewWBService("OWNER", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NJRCI6IjFiMzVmODljLTMyNGYtNGM3OS05NzhhLTkwMmYwODk3Mjc4YiJ9.WeYv1vqA46_9D5up2LRUeSBZCXxSBNcmH8lUhG9Jii0", config)
	jobId := 1
	err = wbService.Parsing(repository, &jobId)
	assert.NoError(t, err, "Error after parsing: %s", err)
	var count int64
	database.Table("stock").Count(&count)
	assert.Equal(t, int64(133), count, "count of stock")
	var stocks []model.StockItem
	database.Where(`"supplier_article"=?`, "ИР060045").Find(&stocks)
	assert.Equal(t, 7, len(stocks), "count of stock")
	stock := stocks[0]
	assert.Equal(t, "ИР060045", *stock.SupplierArticle, "supplier_article")
	assert.Equal(t, "WB", stock.Source, "source")
	assert.Equal(t, float64(636), *stock.Price, "Price")
	lastChangeDate, _ := time.Parse(time.DateOnly, "2023-02-25")
	assert.Equal(t, lastChangeDate, stock.LastChangeDate, "LastChangeDate")
	lastChangeTime, _ := time.Parse(time.TimeOnly, "09:34:40")
	assert.Equal(t, lastChangeTime, stock.LastChangeTime, "LastChangeTime")
	assert.Equal(t, "4620003082726", *stock.Barcode, "barcode")
}
