package main

import (
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

func (r *MessageRepository) Fetch() (result []Message, err error) {
	err = r.DB.Debug().Table("messages").Find(&result).Error
	return
}

func (r *MessageRepository) GetByID(id int) (result Message, err error) {
	err = r.DB.Debug().Table("messages").Where("id = ?", id).First(&result).Error
	return
}

func (r *MessageRepository) GetByUserID(user_id int) (result Message, err error) {
	err = r.DB.Debug().Table("messages").Where("user_id = ?", user_id).First(&result).Error
	return
}

func (r *MessageRepository) GetByRoomID(room_id int) (result Message, err error) {
	err = r.DB.Debug().Table("messages").Where("room_id = ?", room_id).First(&result).Error
	return
}

func (r *MessageRepository) Create(message *Message) (err error) {
	err = r.DB.Debug().Table("messages").Create(message).Error
	return
}
