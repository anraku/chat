package main

import (
	"os"
	"strconv"

	"github.com/anraku/chat/domain"
	"github.com/anraku/chat/trace"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

// Roomは一つのチャットルームを表します
type Room struct {
	ID          int    `gorm:"AUTO_INCREMENT;column:id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	// Forwardは他のクライアントに転送するためのメッセージを保持するチャネルです。
	Forward chan *domain.Message `gorm:"-"`
	// Joinはチャットルームに参加しようとしているクライアントのためのチャネルです。
	Join chan *User `gorm:"-"`
	// Leaveはチャットルームから退室しようとしているクライアントのためのチャネルです
	Leave chan *User `gorm:"-"`
	// Usersには在室しているすべてのクライアントが保持されます。
	Users map[*User]bool `gorm:"-"`
	// Tracerはチャットルーム上で行われた操作のログを受け取ります。
	Tracer trace.Tracer `gorm:"-"`
}

// newRoomはすぐに利用できるチャットルームを生成して返します。
func newRoom(id int) *Room {
	return &Room{
		ID:      id,
		Forward: make(chan *domain.Message),
		Join:    make(chan *User),
		Leave:   make(chan *User),
		Users:   make(map[*User]bool),
		Tracer:  trace.Off(),
	}
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

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  socketBufferSize,
		WriteBufferSize: messageBufferSize,
	}
	rooms = make(map[string]*Room, 1000)
)

// Chat is Handler with WebSocket in chat room
func Chat(c echo.Context) error {
	// WebSocket setting
	roomID := c.Param("id")
	id, err := strconv.Atoi(roomID)
	if err != nil {
		return err
	}
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	// Room setting
	var room *Room
	if _, ok := rooms[roomID]; ok {
		room = rooms[roomID]
	} else {
		room = newRoom(id)
		room.Tracer = trace.New(os.Stdout)
		room.ID = id
		rooms[roomID] = room
		go room.run()
	}

	// get username from session
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	userName := sess.Values["username"].(string)
	client := &User{
		ID:     1,
		Name:   userName,
		Socket: ws,
		Room:   room,
		Send:   make(chan *domain.Message, messageBufferSize),
	}

	// client Join Room
	room.Join <- client
	defer func() { room.Leave <- client }()
	go client.Write()
	client.Read()
	return nil
}
