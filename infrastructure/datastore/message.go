package datastore

import (
	"github.com/anraku/chat/domain/model"
	"github.com/anraku/chat/domain/repository"
	"github.com/jinzhu/gorm"
)

type MessageMySQLRepository struct {
	DB *gorm.DB
}

func NewMessageMySQLRepository(db *gorm.DB) repository.MessageRepository {
	return &MessageMySQLRepository{
		DB: db,
	}
}

func (mr *MessageMySQLRepository) Fetch() (result []model.Message, err error) {
	err = mr.DB.Debug().Table("messages").Find(&result).Error
	return
}

func (mr *MessageMySQLRepository) GetByID(id int) (result model.Message, err error) {
	err = mr.DB.Debug().Table("messages").Where("id = ?", id).First(&result).Error
	return
}

func (mr *MessageMySQLRepository) GetByUserID(user_id int) (result model.Message, err error) {
	err = mr.DB.Debug().Table("messages").Where("user_id = ?", user_id).First(&result).Error
	return
}

func (mr *MessageMySQLRepository) GetByRoomID(room_id int) (result []model.Message, err error) {
	err = mr.DB.Debug().Table("messages").Where("room_id = ?", room_id).Find(&result).Error
	return
}

func (mr *MessageMySQLRepository) Create(message *model.Message) (err error) {
	err = mr.DB.Debug().Table("messages").Create(message).Error
	return
}

func (mr *MessageMySQLRepository) StoreData(m *model.Message) error {
	tx := mr.DB.Begin()
	err := tx.Table("messages").Create(m).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
