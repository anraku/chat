package usecase

import "github.com/anraku/chat/entity"

type MessageInputBoundary interface {
	EnterRoom(*entity.User, *entity.Room)
	GetByRoomID(int) ([]entity.Message, error)
}

type RoomInputBoundary interface {
	Fetch() ([]entity.Room, error)
	Create(entity.Room) error
}

type UserInputBoundary interface {
	SaveUser(interface{}) error
}
