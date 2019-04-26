package service

import (
	"time"

	"github.com/anraku/chat/domain/model"
	"github.com/anraku/chat/domain/repository"
)

type UserService interface {
	Read(*model.User)
	Write(*model.User, repository.MessageRepository)
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (us *userService) Write(u *model.User, mr repository.MessageRepository) {
	for msg := range u.Send {
		if err := u.Socket.WriteJSON(msg); err != nil {
			break
		}

		msg.RoomID = u.Room.ID
		msg.CreatedAt = time.Now()
		//store message
		err := mr.StoreData(msg)
		if err != nil {
			panic(err)
		}
	}
	u.Socket.Close()
}

func (us *userService) Read(u *model.User) {
	for {
		var msg *model.Message
		if err := u.Socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now().Format("2006年01月02日 15:04:05")
			msg.UserName = u.Name
			u.Room.Forward <- msg
		} else {
			break
		}
	}
	u.Socket.Close()
}
