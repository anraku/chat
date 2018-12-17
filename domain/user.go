package domain

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// Userはチャットを行っている1人のユーザーを表します。
type User struct {
	ID   int    `gorm:"AUTO_INCREMENT;column:id"`
	Name string `gorm:"column:name"`
	// socketはこのクライアントのためのWebSocketです。
	Socket *websocket.Conn
	// sendはメッセージが送られるチャネルです。
	Send chan *Message
	// roomはこのクライアントが参加しているチャットルームです。
	Room *Room
	// userDataはユーザーに関する情報を保持します
	UserData *http.Cookie
}

func (user *User) EnterRoom(room *Room, interactor MessageInteractor) {
	// user Join Room
	user.Room = room
	room.Join <- user
	defer func() { room.Leave <- user }()
	// 	go user.Write(interactor)
	// 	user.Read()
}
