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

type messageUsecase struct {
	s  service.MessageService
	mr repository.MessageRepository
}

func NewMessageUsecase(s service.MessageService, mr repository.MessageRepository) MessageUsecase {
	return &messageUsecase{s, mr}
}

func (i *messageUsecase) EnterRoom(user *model.User, room *model.Room) {
	// user join Room
	user.Room = room
	room.Join <- user
	defer func() { room.Leave <- user }()
	go i.s.Write(user)
	i.s.Read(user)
}

func (i *messageUsecase) GetByRoomID(roomID int) (result []model.Message, err error) {
	return i.mr.GetByRoomID(roomID)
}
