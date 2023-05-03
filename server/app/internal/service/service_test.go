package service

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/yoda/app/internal/configuration"
	"github.com/yoda/app/internal/repository"
	"github.com/yoda/common/pkg/model"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var (
	mux     *http.ServeMux
	server  *httptest.Server
	baseUrl string
)

func Test_Main(t *testing.T) {
	fn := setup(t)
	defer fn()

	config := configuration.InitConfig("config.yml")
	config.Database.Dsn = "postgres://root:secret@localhost:5432/pdb"
	config.Database.LoggingLevel = "error"
	config.Ozon.Host = fmt.Sprintf("%s/ozon", baseUrl)
	config.Wb.Host = fmt.Sprintf("%s/wb", baseUrl)
	config.Integration.Host = fmt.Sprintf(`%s/integration`, baseUrl)
	database := repository.InitDatabase(config.Database)
	initTestData(database)

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

func initTestData(db *gorm.DB) {
	query, err := os.ReadFile("./test_data.sql")
	if err != nil {
		panic(err)
	}
	if err := db.Exec(string(query)).Error; err != nil {
		panic(err)
	}
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

func setup(t *testing.T) func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	baseUrl = server.URL
	mockHandlers()

	dockerCompose, err := compose.NewDockerCompose("../../../../docker/docker-compose.yml")
	if err != nil {
		t.Fatalf("compose.NewDockerCompose() error: %v", err)
	}

	t.Cleanup(func() {
		assert.NoError(t, dockerCompose.Down(context.Background(), compose.RemoveOrphans(true), compose.RemoveImagesLocal), "compose.Down()")
	})
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	dockerCompose.Up(ctx, compose.Wait(true))
	time.Sleep(30 * time.Second)
	return func() {
		server.Close()
	}
}

func mockHandlers() {
	orgs, _ := os.ReadFile("../../../../mockdata/dict.organisation.json")
	mux.HandleFunc("/integration/organizations", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(orgs)
	})
	ord, _ := os.ReadFile("../../../../mockdata/wb.order.json")
	mux.HandleFunc("/wb/api/v1/supplier/orders", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(ord)
	})
	sls, _ := os.ReadFile("../../../../mockdata/wb.sale.json")
	mux.HandleFunc("/wb/api/v1/supplier/sales", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(sls)
	})
	stockwb, _ := os.ReadFile("../../../../mockdata/wb.stock.json")
	mux.HandleFunc("/wb/api/v1/supplier/stocks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(stockwb)
	})
	stock, _ := os.ReadFile("../../../../mockdata/ozon.stock.json")
	mux.HandleFunc("/ozon/v2/analytics/stock_on_warehouses", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(stock)
		} else {
			panic("Unsupported method")
		}
	})
	info, _ := os.ReadFile("../../../../mockdata/ozon.info.json")
	mux.HandleFunc("/ozon/v2/product/info/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(info)
		} else {
			panic("Unsupported method")
		}
	})
	prices, _ := os.ReadFile("../../../../mockdata/ozon.prices.json")
	mux.HandleFunc("/ozon/v4/product/info/prices", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(prices)
		} else {
			panic("Unsupported method")
		}
	})
	orders, _ := os.ReadFile("../../../../mockdata/ozon.orders.json")
	mux.HandleFunc("/ozon/v2/posting/fbo/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(orders)
		} else {
			panic("Unsupported method")
		}
	})
	attr, _ := os.ReadFile("../../../../mockdata/ozon.attributes.json")
	mux.HandleFunc("/ozon/v3/products/info/attributes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(attr)
		} else {
			panic("Unsupported method")
		}
	})
}
