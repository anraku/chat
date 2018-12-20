package usecase

import (
	"time"

	"github.com/anraku/chat/entity"
)

type MessageInteractor struct {
	roomRepository    RoomRepository
	messageRepository MessageRepository
}

func NewMessageInteractor(m MessageRepository) *MessageInteractor {
	return &MessageInteractor{
		messageRepository: m,
	}
}

func (i *MessageInteractor) write(user *entity.User) {
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

func (i *MessageInteractor) read(user *entity.User) {
	for {
		var msg *entity.Message
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

func (i *MessageInteractor) EnterRoom(user *entity.User, room *entity.Room) {
	// user join Room
	user.Room = room
	room.Join <- user
	defer func() { room.Leave <- user }()
	go i.write(user)
	i.read(user)
}

func (i *MessageInteractor) GetByRoomID(roomID int) (result []entity.Message, err error) {
	result, err = i.messageRepository.GetByRoomID(roomID)
	return
}
