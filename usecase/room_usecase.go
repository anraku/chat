package usecase

import (
	"github.com/anraku/chat/domain/model"
	"github.com/anraku/chat/domain/repository"
)

type RoomUsecase interface {
	Fetch() ([]model.Room, error)
	Create(model.Room) error
}

type roomUsecase struct {
	rr repository.RoomRepository
}

func NewRoomUsecase(rr repository.RoomRepository) RoomUsecase {
	return &roomUsecase{rr}
}

func (ur *roomUsecase) Fetch() (rooms []model.Room, err error) {
	return ur.rr.Fetch()
}

func (ur *roomUsecase) Create(room model.Room) (err error) {
	return ur.rr.Create(room)
}
