package interfaces

import "github.com/anraku/chat/entity"

type MessageInteractor interface {
	EnterRoom(*entity.User, *entity.Room)
	GetByRoomID(int) ([]entity.Message, error)
}

type RoomInteractor interface {
	Fetch() ([]entity.Room, error)
	Create(entity.Room) error
}

type UserInteractor interface {
	StoreData(*entity.Message) error
}
