package repository

import (
	"github.com/anraku/chat/domain"
	"github.com/jinzhu/gorm"
)

type RoomRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{
		DB: db,
	}
}

const table = "rooms"

func (r *RoomRepository) Fetch() (result []domain.Room, err error) {
	err = r.DB.Debug().Table(table).Find(&result).Error
	return
}

func (r *RoomRepository) GetByID(id int) (result domain.Room, err error) {
	err = r.DB.Debug().Table(table).Where("id = ?", id).First(&result).Error
	return
}
