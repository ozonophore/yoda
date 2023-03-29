package dao

import (
	"errors"
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
