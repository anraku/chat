package usecase

import (
	"github.com/anraku/chat/domain/model"
)

type RoomUsecase interface {
	Fetch() ([]model.Room, error)
	Create(model.Room) error
}

type RoomInteractor struct {
	roomRepository    RoomRepository
	messageRepository MessageRepository
}

func NewRoomInteractor(rr RoomRepository, mr MessageRepository) RoomUsecase {
	return &RoomInteractor{
		roomRepository:    rr,
		messageRepository: mr,
	}
}

func (interactor *RoomInteractor) Fetch() (rooms []model.Room, err error) {
	rooms, err = interactor.roomRepository.Fetch()
	return
}

func (interactor *RoomInteractor) Create(room model.Room) (err error) {
	err = interactor.roomRepository.Create(room)
	return
}
