package service

import (
	"os"

	"github.com/anraku/chat/domain/model"
	"github.com/anraku/chat/infrastructure/trace"
)

type RoomService interface {
	GetRoom(id int) *model.Room
}

type roomService struct{}

func NewRoomService() RoomService {
	return &roomService{}
}

var rooms = make(map[int]*model.Room, 1000)

func (rs *roomService) GetRoom(id int) *model.Room {
	// Room setting
	var room *model.Room
	if _, ok := rooms[id]; ok {
		room = rooms[id]
	} else {
		room = &model.Room{
			ID:      id,
			Forward: make(chan *model.Message),
			Join:    make(chan *model.User),
			Leave:   make(chan *model.User),
			Users:   make(map[*model.User]bool),
			Tracer:  trace.New(os.Stdout),
		}
		rooms[id] = room
		go rs.run(room)
	}
	return room
}

func (rs *roomService) run(r *model.Room) {
	for {
		select {
		case client := <-r.Join:
			// 参加
			r.Users[client] = true
			r.Tracer.Trace("新しいクライアントが参加しました")
		case client := <-r.Leave:
			// 退室
			delete(r.Users, client)
			close(client.Send)
			r.Tracer.Trace("クライアントが退室しました")
		case msg := <-r.Forward:
			r.Tracer.Trace("メッセージを受信しました: ", msg.Message)
			// すべてのクライアントにメッセージを転送
			for client := range r.Users {
				select {
				case client.Send <- msg:
					// メッセージを送信
					r.Tracer.Trace(" -- クライアントに送信されました")
				default:
					// 送信に失敗
					delete(r.Users, client)
					close(client.Send)
					r.Tracer.Trace(" -- 送信に失敗しました。クライアントをクリーンアップします")
				}
			}
		}
	}
}
