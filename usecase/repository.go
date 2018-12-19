package usecase

import "github.com/anraku/chat/entity"

type RoomRepository interface {
	Fetch() (rooms []entity.Room, err error)
	GetByID(id int) (result entity.Room, err error)
	Create(room entity.Room) (err error)
}

type UserRepository interface {
	Create(*entity.Message) error
}

type MessageRepository interface {
	Fetch() (result []entity.Message, err error)
	GetByID(id int) (result entity.Message, err error)
	GetByUserID(user_id int) (result entity.Message, err error)
	GetByRoomID(room_id int) (result []entity.Message, err error)
	Create(message *entity.Message) (err error)
	StoreData(m *entity.Message) error
}
