package usecase

import "github.com/anraku/chat/domain/model"

type RoomRepository interface {
	Fetch() (rooms []model.Room, err error)
	GetByID(id int) (result model.Room, err error)
	Create(room model.Room) (err error)
}

type UserRepository interface {
	Create(interface{}) error
}

type MessageRepository interface {
	Fetch() (result []model.Message, err error)
	GetByID(id int) (result model.Message, err error)
	GetByUserID(user_id int) (result model.Message, err error)
	GetByRoomID(room_id int) (result []model.Message, err error)
	Create(message *model.Message) (err error)
	StoreData(m *model.Message) error
}
