package domain

import (
	"net/http"
	"time"

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

func (user *User) Read() {
	for {
		var msg *Message
		if err := user.Socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now().Format("2006年01月02日 15:04:05")
			msg.UserName = user.Name // retrieve username from session
			user.Room.Forward <- msg
		} else {
			break
		}
	}
	user.Socket.Close()
}

func (user *User) Write(interactor MessageInteractor) {
	for msg := range user.Send {
		if err := user.Socket.WriteJSON(msg); err != nil {
			break
		}
		msg.RoomID = user.Room.ID
		msg.CreatedAt = time.Now()
		//store message
		err := interactor.StoreData(msg)
		if err != nil {
			panic(err)
		}
	}
	user.Socket.Close()
}

func (user *User) EnterRoom(room *Room, interactor MessageInteractor) {
	// user Join Room
	user.Room = room
	room.Join <- user
	defer func() { room.Leave <- user }()
	go user.Write(interactor)
	user.Read()
}
