package repository

import "github.com/anraku/chat/model"

type MessageRepository interface {
	Fetch() (result []model.Message, err error)
	GetByID(id int) (result model.Message, err error)
	GetByUserID(user_id int) (result model.Message, err error)
	GetByRoomID(room_id int) (result []model.Message, err error)
	Create(message *model.Message) (err error)
	StoreData(m *model.Message) error
}
