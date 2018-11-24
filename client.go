package main

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Clientはチャットを行っている1人のユーザーを表します。
type Client struct {
	ID   int
	Name string
	// socketはこのクライアントのためのWebSocketです。
	Socket *websocket.Conn
	// sendはメッセージが送られるチャネルです。
	Send chan *Message
	// roomはこのクライアントが参加しているチャットルームです。
	Room *Room
	// userDataはユーザーに関する情報を保持します
	UserData *http.Cookie
}

func (c *Client) Read() {
	for {
		var msg *Message
		if err := c.Socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.UserName = c.UserData.Value // retrieve username from cookie
			c.Room.Forward <- msg
		} else {
			break
		}
	}
	c.Socket.Close()
}

func (c *Client) Write() {
	for msg := range c.Send {
		if err := c.Socket.WriteJSON(msg); err != nil {
			break
		}
		msg.RoomID = c.Room.ID
		msg.UserID = c.ID
		//store message
		err := storeData(msg)
		if err != nil {
			panic(err)
		}
	}
	c.Socket.Close()
}

func storeData(m *Message) error {
	tx := DB.Begin()
	err := NewMessageRepository(tx).Create(m)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
