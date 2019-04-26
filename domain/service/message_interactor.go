package service

import (
	"time"

	"github.com/anraku/chat/domain/model"
	"github.com/anraku/chat/domain/repository"
)

type MessageService interface {
	Write(*model.User)
	Read(*model.User)
}

type messageService struct {
	mr repository.MessageRepository
}

func NewMessageService(mr repository.MessageRepository) MessageService {
	return &messageService{mr}
}

func (s *messageService) Write(user *model.User) {
	for msg := range user.Send {
		if err := user.Socket.WriteJSON(msg); err != nil {
			break
		}
		msg.RoomID = user.Room.ID
		msg.CreatedAt = time.Now()
		//store message
		err := s.mr.StoreData(msg)
		if err != nil {
			panic(err)
		}
	}
	user.Socket.Close()
}

func (s *messageService) Read(user *model.User) {
	for {
		var msg *model.Message
		if err := user.Socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now().Format("2006年01月02日 15:04:05")
			msg.UserName = user.Name
			user.Room.Forward <- msg
		} else {
			break
		}
	}
	user.Socket.Close()
}
