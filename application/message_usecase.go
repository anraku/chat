package usecase

import (
	"time"

	"github.com/anraku/chat/entity"
	"github.com/anraku/chat/model/model"
	"github.com/anraku/chat/model/repository"
)

type MessageUsecase interface {
	EnterRoom(*entity.User, *entity.Room)
	GetByRoomID(int) ([]entity.Message, error)
}

type MessageInteractor struct {
	roomRepository    repository.RoomRepository
	messageRepository repository.MessageRepository
}

func NewMessageInteractor(m repository.MessageRepository) *MessageInteractor {
	return &MessageInteractor{
		messageRepository: m,
	}
}

func (i *MessageInteractor) write(user *model.User) {
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

func (i *MessageInteractor) read(user *model.User) {
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

func (i *MessageInteractor) EnterRoom(user *model.User, room *model.Room) {
	// user join Room
	user.Room = room
	room.Join <- user
	defer func() { room.Leave <- user }()
	go i.write(user)
	i.read(user)
}

func (i *MessageInteractor) GetByRoomID(roomID int) (result []model.Message, err error) {
	result, err = i.messageRepository.GetByRoomID(roomID)
	return
}
