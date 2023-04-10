package dao

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/yoda/common/pkg/model"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	database *gorm.DB
}

var dao *Repository

func NewRepositoryDAO(database *gorm.DB) *Repository {
	if dao == nil {
		dao = &Repository{
			database: database,
		}
	}
	return dao
}

func (p *Repository) CreateTlgMessage(ChatID int64, ChatType string, MessageID int, Message *string) (*model.TlgMessage, error) {
	model := model.TlgMessage{
		ChatID:      ChatID,
		ChatType:    ChatType,
		MessageID:   MessageID,
		Message:     Message,
		MessageDate: time.Now(),
		Status:      model.TlgMessageStatusCreated,
	}
	if err := p.database.Create(&model).Error; err != nil {
		panic(err)
	}
	return &model, nil
}

func (p *Repository) SetTlgMessageStatus(id int64, status string) error {
	if err := p.database.Model(&model.TlgMessage{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func GetTlgMessageById(id int64) (*model.TlgMessage, error) {
	var model model.TlgMessage
	if err := dao.database.Where(`"id" = ?`, id).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model, nil
}

func ExistsTlgEventByChatIdAndEvent(chatId int64, dataType string) bool {
	var count int64
	dao.database.Model(&model.TlgEvent{}).Where(`"chat_id" = ? and "data_type" = ?`, chatId, dataType).Count(&count)
	return count > 0
}

func CreateTlgEvent(chatId int64, dataType string) error {
	model := model.TlgEvent{
		ChatID:   chatId,
		DataType: dataType,
	}
	if err := dao.database.Create(&model).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTlgEvent(chatId int64, dataType string) error {
	if err := dao.database.Where(`"chat_id" = ? and "data_type" = ?`, chatId, dataType).Delete(&model.TlgEvent{}).Error; err != nil {
		return err
	}
	return nil
}

func SaveOwner(o model.Owner, oz model.OwnerMarketplace, wb model.OwnerMarketplace) error {
	tx := dao.database.Begin()
	defer func() {
		err := tx.Error
		if err == nil {
			err := tx.Commit().Error
			if err != nil {
				logrus.Error(err)
			}
		} else {
			err := tx.Rollback()
			logrus.Error(err)
		}
	}()

	var count int64
	tx.Model(&o).Where(`"code" = ?`, o.Code).Count(&count)
	var err error
	if count == 0 {
		err = tx.Create(&o).Error
	} else {
		err = tx.Model(&o).Omit("create_date").Updates(o).Error
	}
	if err != nil {
		return err
	}
	err = tx.Save(&oz).Error
	if err != nil {
		return err
	}
	err = tx.Save(&wb).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOwnerByCode(code string) (*model.Owner, error) {
	var model model.Owner
	if err := dao.database.Where(`"code" = ?`, code).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model, nil
}

func GetOwnerMarketplaceByOwnerCodeAndSource(ownerCode string, source string) (*model.OwnerMarketplace, error) {
	var model model.OwnerMarketplace
	if err := dao.database.Where(`"owner_code" = ? and "source" = ?`, ownerCode, source).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model, nil
}

func GetJobs() ([]model.Job, error) {
	var models []model.Job
	if err := dao.database.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func GetJobById(id int64) (*model.Job, error) {
	var model model.Job
	if err := dao.database.Where(`"id" = ?`, id).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &model, nil
}

func GetOwners() ([]model.Owner, error) {
	var models []model.Owner
	if err := dao.database.Order(`"create_date" desc`).Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func GetMarketplacesByOwners(owners []model.Owner) (*[]model.OwnerMarketplace, error) {
	var models []model.OwnerMarketplace
	ids := make([]string, len(owners))
	for i, owner := range owners {
		ids[i] = owner.Code
	}
	if err := dao.database.Where(`"owner_code" in ?`, ids).Find(&models).Error; err != nil {
		return nil, err
	}
	return &models, nil
}
