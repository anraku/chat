package usecase

import (
	"github.com/anraku/chat/domain/model"
	"github.com/anraku/chat/domain/repository"
)

type MessageUsecase interface {
	GetByRoomID(int) ([]model.Message, error)
}

type MessageInteractor struct {
	mr repository.MessageRepository
}

func NewMessageInteractor(m repository.MessageRepository) MessageUsecase {
	return &MessageInteractor{
		mr: m,
	}
}

func (mi *MessageInteractor) GetByRoomID(roomID int) (result []model.Message, err error) {
	result, err = mi.mr.GetMessagesByRoomID(roomID)
	return
}
