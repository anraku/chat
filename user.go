package main

import (
	"net/http"
	"time"

	"github.com/anraku/chat/domain"
	"github.com/gorilla/websocket"
)

// Userはチャットを行っている1人のユーザーを表します。
type User struct {
	ID   int    `gorm:"AUTO_INCREMENT;column:id"`
	Name string `gorm:"column:name"`
	// socketはこのクライアントのためのWebSocketです。
	Socket *websocket.Conn
	// sendはメッセージが送られるチャネルです。
	Send chan *domain.Message
	// roomはこのクライアントが参加しているチャットルームです。
	Room *Room
	// userDataはユーザーに関する情報を保持します
	UserData *http.Cookie
}

func (c *User) Read() {
	for {
		var msg *domain.Message
		if err := c.Socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now().Format("2006年01月02日 15:04:05")
			msg.UserName = c.Name // retrieve username from session
			c.Room.Forward <- msg
		} else {
			break
		}
	}
	c.Socket.Close()
}

func (c *User) Write() {
	for msg := range c.Send {
		if err := c.Socket.WriteJSON(msg); err != nil {
			break
		}
		msg.RoomID = c.Room.ID
		msg.CreatedAt = time.Now()
		//store message
		err := storeData(msg)
		if err != nil {
			panic(err)
		}
	}
	c.Socket.Close()
}

func storeData(m *domain.Message) error {
	tx := DB.Begin()
	err := NewMessageRepository(tx).Create(m)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
