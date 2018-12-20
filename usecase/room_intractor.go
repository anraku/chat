package usecase

import (
	"github.com/anraku/chat/entity"
)

type RoomInteractor struct {
	roomRepository    RoomRepository
	messageRepository MessageRepository
}

func NewRoomInteractor(rr RoomRepository, mr MessageRepository) *RoomInteractor {
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
