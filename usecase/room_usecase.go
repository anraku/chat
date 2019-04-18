package usecase

import (
	"github.com/anraku/chat/domain/model"
	"github.com/anraku/chat/domain/repository"
)

type RoomUsecase interface {
	Fetch() ([]model.Room, error)
	Create(model.Room) error
	GetMessage(int) ([]model.Message, error)
}

type RoomInteractor struct {
	us service.UserService
	mr repository.MessageRepository
	
	rs service.RoomService
	rr repository.RoomRepository
}

func NewRoomInteractor(rs RoomService) RoomUsecase {
	return &RoomInteractor{
		rs: rs,
	}
}

func (ri *RoomInteractor) Fetch() (rooms []model.Room, err error) {
	rooms, err = ri.rr.Fetch()
	return
}

func (ri *RoomInteractor) GetMessages(roomID int) (result []model.Message, err error) {
	result, err = mi.mr.GetMessagesByRoomID(roomID)
	return
}

func (ri *RoomInteractor) Create(room model.Room) (err error) {
	err = ri.rr.Create(room)
	return
}

func (ri *RoomInteractor) EnterRoom(user *model.User, room *model.Room) {
	// user join Room
	user.Room = room
	room.Join <- user
	defer func() { room.Leave <- user }()
	go ri.us.Write(user, mi.mr)
	ri.us.Read(user, mi.mr)
}

