package service

import (
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/storage"
	"github.com/yoda/common/pkg/model"
)

func CreateLogLoad(transactionId int64, owner string, source string) error {
	return storage.CreateOrUpdateLogLoad(&model.LogLoad{
		TransactionID: transactionId,
		OwnerCode:     owner,
		Source:        source,
		Status:        "BEGIN",
	})
}

func CompleteLogLoad(transactionId int64, owner string, source string) {
	err := storage.CreateOrUpdateLogLoad(&model.LogLoad{
		TransactionID: transactionId,
		OwnerCode:     owner,
		Source:        source,
		Status:        "COMPLETED",
	})
	if err != nil {
		logrus.Panic(err)
	}
}

func ErrorLogLoad(transactionId int64, owner string, source string, e error) {
	msg := e.Error()
	err := storage.CreateOrUpdateLogLoad(&model.LogLoad{
		TransactionID: transactionId,
		OwnerCode:     owner,
		Source:        source,
		Description:   &msg,
		Status:        "ERROR",
	})
	if err != nil {
		logrus.Panic(err)
	}
}
