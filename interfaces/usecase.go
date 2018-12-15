package interfaces

import "github.com/anraku/chat/domain"

type MessageInteractor interface {
	GetByRoomID(int) ([]domain.Message, error)
	StoreData(*domain.Message) error
}
