package service

import (
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/pkg/repository"
	"github.com/yoda/common/pkg/model"
)

func CreateLogLoad(transactionId int64, owner string, source string) error {
	return repository.CreateOrUpdateLogLoad(&model.LogLoad{
		TransactionID: transactionId,
		OwnerCode:     owner,
		Source:        source,
		Status:        "BEGIN",
	})
}

func CompleteLogLoad(transactionId int64, owner string, source string) {
	err := repository.CreateOrUpdateLogLoad(&model.LogLoad{
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
	err := repository.CreateOrUpdateLogLoad(&model.LogLoad{
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
