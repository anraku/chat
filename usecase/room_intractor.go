package usecase

import (
	"github.com/anraku/chat/entity"
	"github.com/anraku/chat/interfaces"
)

type RoomInteractor struct {
	roomRepository    interfaces.RoomRepository
	RoomPresenter     RoomOutputBoundary
	messageRepository interfaces.MessageRepository
}

func NewRoomInteractor(rr interfaces.RoomRepository, mr interfaces.MessageRepository) *RoomInteractor {
	return &RoomInteractor{
		roomRepository:    rr,
		messageRepository: mr,
	}
}

func (interactor *RoomInteractor) Fetch() (rooms []entity.Room, err error) {
	rooms, err = interactor.roomRepository.Fetch()
	return
}

func (interactor *RoomInteractor) Create(room entity.Room) (err error) {
	err = interactor.roomRepository.Create(room)
	return
}
