package main

import (
	"net/http"
	"time"

	"github.com/anraku/chat/domain"
	"github.com/gorilla/websocket"
)

// clientはチャットを行っている1人のユーザーを表します。
type client struct {
	// socketはこのクライアントのためのWebSocketです。
	socket *websocket.Conn
	// sendはメッセージが送られるチャネルです。
	send chan *domain.Message
	// roomはこのクライアントが参加しているチャットルームです。
	room *room
	// userDataはユーザーに関する情報を保持します
	userData *http.Cookie
}

func (c *client) read() {
	for {
		var msg *domain.Message
		if err := c.socket.ReadJSON(&msg); err == nil {
			//storeData(msg)
			msg.When = time.Now()
			msg.UserName = c.userData.Value // retrieve username from cookie
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}

func storeData(m *domain.Message) error {
	return nil
}
