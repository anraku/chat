package usecase

import (
	"github.com/anraku/chat/domain"
	"github.com/anraku/chat/interfaces"
)

type MessageInteractor struct {
	roomRepository    interfaces.RoomRepository
	messageRepository interfaces.MessageRepository
}

func NewMessageInteractor(m interfaces.MessageRepository) *MessageInteractor {
	return &MessageInteractor{
		messageRepository: m,
	}
}

func (i *MessageInteractor) GetByRoomID(roomID int) (result []domain.Message, err error) {
	result, err = i.messageRepository.GetByRoomID(roomID)
	return
}
