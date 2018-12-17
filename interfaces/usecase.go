package interfaces

import "github.com/anraku/chat/domain"

type MessageInteractor interface {
	EnterRoom(*domain.User, *domain.Room)
	GetByRoomID(int) ([]domain.Message, error)
}

type RoomInteractor interface {
	Fetch() ([]domain.Room, error)
	Create(domain.Room) error
}

type UserInteractor interface {
	StoreData(*domain.Message) error
}
