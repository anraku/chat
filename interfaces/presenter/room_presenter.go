package presenter

import (
	"github.com/anraku/chat/entity"
	"github.com/anraku/chat/interfaces"
	"github.com/anraku/chat/usecase"
)

type RoomPresenter struct{}

func NewRoomPresenter() usecase.RoomOutputBoundary {
	return &RoomPresenter{}
}

func (p *RoomPresenter) CreateIndexData(c interfaces.Context, rooms []entity.Room) map[string]interface{} {
	// get username from context
	var username string
	if v, ok := c.Get("username").(string); ok {
		username = v
	} else {
		username = ""
	}
	data := map[string]interface{}{
		"username": username,
		"rooms":    rooms,
	}
	return data
}
