package datastore

import (
	"github.com/anraku/chat/domain/model"
	"github.com/anraku/chat/domain/repository"
	"github.com/jinzhu/gorm"
)

type RoomMySQLRepository struct {
	DB *gorm.DB
}

func NewRoomMySQLRepository(db *gorm.DB) repository.RoomRepository {
	return &RoomMySQLRepository{
		DB: db,
	}
}

func (rr *RoomMySQLRepository) Fetch() (result []model.Room, err error) {
	err = rr.DB.Table("rooms").Find(&result).Error
	return
}

func (rr *RoomMySQLRepository) GetByID(id int) (result model.Room, err error) {
	err = rr.DB.Table("rooms").Where("id = ?", id).First(&result).Error
	return
}

func (rr *RoomMySQLRepository) Create(room model.Room) (err error) {
	err = rr.DB.Table("rooms").Create(&room).Error
	return
}
