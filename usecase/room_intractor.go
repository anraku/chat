package usecase

import (
	"github.com/anraku/chat/domain"
	"github.com/anraku/chat/interfaces"
)

type RoomInteractor struct {
	roomRepository    interfaces.RoomRepository
	messageRepository interfaces.MessageRepository
}

func NewRoomInteractor(r interfaces.RoomRepository, m interfaces.MessageRepository) interfaces.RoomInteractor {
	return &RoomInteractor{
		roomRepository:    r,
		messageRepository: m,
	}
}

func (interactor *RoomInteractor) Fetch() (rooms []domain.Room, err error) {
	rooms, err = interactor.roomRepository.Fetch()
	return
}

func (interactor *RoomInteractor) Create(room domain.Room) (err error) {
	err = interactor.roomRepository.Create(room)
	return
}
