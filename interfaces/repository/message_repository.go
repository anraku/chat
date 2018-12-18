package repository

import (
	"github.com/anraku/chat/entity"
	"github.com/jinzhu/gorm"
)

type MessageRepository struct {
	DB *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{
		DB: db,
	}
}

func (r *MessageRepository) Fetch() (result []entity.Message, err error) {
	err = r.DB.Debug().Table("messages").Find(&result).Error
	return
}

func (r *MessageRepository) GetByID(id int) (result entity.Message, err error) {
	err = r.DB.Debug().Table("messages").Where("id = ?", id).First(&result).Error
	return
}

func (r *MessageRepository) GetByUserID(user_id int) (result entity.Message, err error) {
	err = r.DB.Debug().Table("messages").Where("user_id = ?", user_id).First(&result).Error
	return
}

func (r *MessageRepository) GetByRoomID(room_id int) (result []entity.Message, err error) {
	err = r.DB.Debug().Table("messages").Where("room_id = ?", room_id).Find(&result).Error
	return
}

func (r *MessageRepository) Create(message *entity.Message) (err error) {
	err = r.DB.Debug().Table("messages").Create(message).Error
	return
}

func (r *MessageRepository) StoreData(m *entity.Message) error {
	tx := r.DB.Begin()
	err := tx.Table("messages").Create(m).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
