package usecase

import (
	"github.com/anraku/chat/domain/model"
	"github.com/anraku/chat/domain/repository"
	"github.com/anraku/chat/domain/service"
)

type MessageUsecase interface {
	EnterRoom(*model.User, *model.Room)
	GetByRoomID(int) ([]model.Message, error)
}

type MessageInteractor struct {
	s  service.MessageService
	mr repository.MessageRepository
}

func NewMessageInteractor(s service.MessageService, mr repository.MessageRepository) MessageUsecase {
	return &MessageInteractor{s, mr}
}

func (i *MessageInteractor) EnterRoom(user *model.User, room *model.Room) {
	// user join Room
	user.Room = room
	room.Join <- user
	defer func() { room.Leave <- user }()
	go i.s.Write(user)
	i.s.Read(user)
}

func (i *MessageInteractor) GetByRoomID(roomID int) (result []model.Message, err error) {
	return i.mr.GetByRoomID(roomID)
}
