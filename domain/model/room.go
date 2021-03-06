package model

import (
	"os"

	"github.com/anraku/chat/interfaces/trace"
)

// Roomは一つのチャットルームを表します
type Room struct {
	ID          int    `gorm:"AUTO_INCREMENT;column:id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	// Forwardは他のクライアントに転送するためのメッセージを保持するチャネルです。
	Forward chan *Message `gorm:"-"`
	// Joinはチャットルームに参加しようとしているクライアントのためのチャネルです。
	Join chan *User `gorm:"-"`
	// Leaveはチャットルームから退室しようとしているクライアントのためのチャネルです
	Leave chan *User `gorm:"-"`
	// Usersには在室しているすべてのクライアントが保持されます。
	Users map[*User]bool `gorm:"-"`
	// Tracerはチャットルーム上で行われた操作のログを受け取ります。
	Tracer trace.Tracer `gorm:"-"`
}

var rooms = make(map[int]*Room, 1000)

func GetRoom(id int) *Room {
	// Room setting
	var room *Room
	if _, ok := rooms[id]; ok {
		room = rooms[id]
	} else {
		room = &Room{
			ID:      id,
			Forward: make(chan *Message),
			Join:    make(chan *User),
			Leave:   make(chan *User),
			Users:   make(map[*User]bool),
			Tracer:  trace.New(os.Stdout),
		}
		rooms[id] = room
		go room.run()
	}
	return room
}

func (r *Room) run() {
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
