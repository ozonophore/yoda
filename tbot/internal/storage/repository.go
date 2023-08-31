package storage

import (
	"github.com/sirupsen/logrus"
	"github.com/yoda/tnot/internal/configuration"
	"github.com/yoda/tnot/internal/storage/notification"
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

func (r *Repository) GetReportByProduct(date time.Time) (*[]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := r.db.Raw(`select 
    t.report_date as total_report_date,
	t.source as total_source,
	t.item_id as total_item_id,
	t.item_name as total_item_name,
	t.def30 as total_def30,
	t.def5 as total_def5,
	t.avg_price as total_avg_price,
	t.min_price as total_min_price,
	t.max_price as total_max_price,
	t.quantity30 as total_quantity30,
	t.quantity5 as total_quantity5,
	t.order_by_day30 as total_order_by_day30,
	t.forecast_order30 as total_forecast_order30,
	t.order_by_day5 as total_order_by_day5,
	t.forecast_order5 as total_forecast_order5,
	t.quantity as total_quantity,
	t.is_excluded as total_is_excluded,
	t.quantity1c as total_quantity1c,
	t.quantity5_week_ago as total_quantity5_week_ago,
	t.turnover30 as total_turnover30,
	t.turnover5 as total_turnover5,
	t.stock_total as total_stock_total,
	t.stock1c_percent as total_stock1c_percent,
	t.segment,
	t.brand,
	t.retail_price,
	wb.def30 as wb_def30,
	wb.def5 as wb_def5,
	wb.avg_price as wb_avg_price,
	wb.min_price as wb_min_price,
	wb.max_price as wb_max_price,
	wb.quantity30 as wb_quantity30,
	wb.quantity5 as wb_quantity5,
	wb.order_by_day30 as wb_order_by_day30,
	wb.forecast_order30 as wb_forecast_order30,
	wb.order_by_day5 as wb_order_by_day5,
	wb.forecast_order5 as wb_forecast_order5,
	wb.quantity as wb_quantity,
	wb.is_excluded as wb_is_excluded,
	wb.quantity1c as wb_quantity1c,
	wb.quantity5_week_ago as wb_quantity5_week_ago,
	wb.turnover30 as wb_turnover30,
	wb.turnover5 as wb_turnover5,
	wb.stock_total as wb_stock_total,
	wb.stock1c_percent as wb_stock1c_percent,
	ozon.def30 as ozon_def30,
	ozon.def5 as ozon_def5,
	ozon.avg_price as ozon_avg_price,
	ozon.min_price as ozon_min_price,
	ozon.max_price as ozon_max_price,
	ozon.quantity30 as ozon_quantity30,
	ozon.quantity5 as ozon_quantity5,
	ozon.order_by_day30 as ozon_order_by_day30,
	ozon.forecast_order30 as ozon_forecast_order30,
	ozon.order_by_day5 as ozon_order_by_day5,
	ozon.forecast_order5 as ozon_forecast_order5,
	ozon.quantity as ozon_quantity,
	ozon.is_excluded as ozon_is_excluded,
	ozon.quantity1c as ozon_quantity1c,
	ozon.quantity5_week_ago as ozon_quantity5_week_ago,
	ozon.turnover30 as ozon_turnover30,
	ozon.turnover5 as ozon_turnover5,
	ozon.stock_total as ozon_stock_total,
	ozon.stock1c_percent as ozon_stock1c_percent
    from dl.report_by_product t
         left outer join dl.report_by_product wb on t.item_id = wb.item_id and wb.source = 'WB' and wb.report_date = t.report_date
         left outer join dl.report_by_product ozon on t.item_id = ozon.item_id and ozon.source = 'OZON' and ozon.report_date = t.report_date
         where t.report_date = ? and t.source=?`, date.Format(time.DateOnly), "TOTAL").Find(&results).Error
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (r *Repository) GetReport(date time.Time) (*[]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := r.db.Table("dl.sales_stock").Where(`"report_date" = ?`, date.Format(time.DateOnly)).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (r *Repository) GetReportByCluster(date time.Time) (*[]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := r.db.Table("dl.report_by_cluster").Where(`"report_date" = ?`, date.Format(time.DateOnly)).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (r *Repository) GetReportByItem(date time.Time) (*[]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := r.db.Table("dl.report_by_item").Where(`"report_date" = ?`, date.Format(time.DateOnly)).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return &results, nil
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
