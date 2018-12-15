package interfaces

import "github.com/anraku/chat/domain"

type RoomRepository interface {
	Fetch() (rooms []domain.Room, err error)
	GetByID(id int) (result domain.Room, err error)
	Create(room domain.Room) (err error)
}

type UserRepository interface{}

type MessageRepository interface {
	Fetch() (result []domain.Message, err error)
	GetByID(id int) (result domain.Message, err error)
	GetByUserID(user_id int) (result domain.Message, err error)
	GetByRoomID(room_id int) (result []domain.Message, err error)
	Create(message *domain.Message) (err error)
	StoreData(m *domain.Message) error
}
