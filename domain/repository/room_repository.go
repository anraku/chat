package repository

import "github.com/anraku/chat/model"

type RoomRepository interface {
	Fetch() (rooms []model.Room, err error)
	GetByID(id int) (result model.Room, err error)
	Create(room model.Room) (err error)
}
