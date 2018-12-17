package usecase

import (
	"time"

	"github.com/anraku/chat/domain"
	"github.com/anraku/chat/interfaces"
)

type MessageInteractor struct {
	roomRepository    interfaces.RoomRepository
	messageRepository interfaces.MessageRepository
}

func NewMessageInteractor(m interfaces.MessageRepository) interfaces.MessageInteractor {
	return &MessageInteractor{
		messageRepository: m,
	}
}

func (i *MessageInteractor) Write(user *domain.User) {
	for msg := range user.Send {
		if err := user.Socket.WriteJSON(msg); err != nil {
			break
		}
		msg.RoomID = user.Room.ID
		msg.CreatedAt = time.Now()
		//store message
		err := i.messageRepository.StoreData(msg)
		if err != nil {
			panic(err)
		}
	}
	user.Socket.Close()
}

func (i *MessageInteractor) Read(user *domain.User) {
	for {
		var msg *domain.Message
		if err := user.Socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now().Format("2006年01月02日 15:04:05")
			msg.UserName = user.Name // retrieve username from session
			user.Room.Forward <- msg
		} else {
			break
		}
	}
	user.Socket.Close()
}

func (i *MessageInteractor) GetByRoomID(roomID int) (result []domain.Message, err error) {
	result, err = i.messageRepository.GetByRoomID(roomID)
	return
}

func (i *MessageInteractor) StoreData(m *domain.Message) error {
	return i.messageRepository.StoreData(m)
}
