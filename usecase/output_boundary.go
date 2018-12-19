package usecase

import (
	"github.com/anraku/chat/entity"
	"github.com/anraku/chat/interfaces"
)

type UserOutputBoundary interface {
}

type RoomOutputBoundary interface {
	CreateIndexData(interfaces.Context, []entity.Room) map[string]interface{}
}
