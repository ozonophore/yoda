package storage

import (
	"github.com/yoda/common/pkg/model"
)

func SaveReportDetail(items *[]model.ReportDetailByPeriod, batchSize int) error {
	initIfError()
	return repository.db.CreateInBatches(items, batchSize).Error
}
