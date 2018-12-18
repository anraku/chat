package repository

import (
	"github.com/anraku/chat/entity"
	"github.com/jinzhu/gorm"
)

type RoomRepository struct {
	DB *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{
		DB: db,
	}
}

func (r *RoomRepository) Fetch() (result []entity.Room, err error) {
	err = r.DB.Debug().Table("rooms").Find(&result).Error
	return
}

func (r *RoomRepository) GetByID(id int) (result entity.Room, err error) {
	err = r.DB.Debug().Table("rooms").Where("id = ?", id).First(&result).Error
	return
}

func (r *RoomRepository) Create(room entity.Room) (err error) {
	err = r.DB.Debug().Table("rooms").Create(&room).Error
	return
}