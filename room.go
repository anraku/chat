package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/anraku/chat/trace"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

// Roomは一つのチャットルームを表します
type Room struct {
	ID          int    `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	// Forwardは他のクライアントに転送するためのメッセージを保持するチャネルです。
	Forward chan *Message
	// Joinはチャットルームに参加しようとしているクライアントのためのチャネルです。
	Join chan *Client
	// Leaveはチャットルームから退室しようとしているクライアントのためのチャネルです
	Leave chan *Client
	// Clientsには在室しているすべてのクライアントが保持されます。
	Clients map[*Client]bool
	// Tracerはチャットルーム上で行われた操作のログを受け取ります。
	Tracer trace.Tracer
}

// newRoomはすぐに利用できるチャットルームを生成して返します。
func newRoom(id int) *Room {
	return &Room{
		ID:      id,
		Forward: make(chan *Message),
		Join:    make(chan *Client),
		Leave:   make(chan *Client),
		Clients: make(map[*Client]bool),
		Tracer:  trace.Off(),
	}
}

func (r *Room) run() {
	for {
		select {
		case client := <-r.Join:
			// 参加
			r.Clients[client] = true
			r.Tracer.Trace("新しいクライアントが参加しました")
		case client := <-r.Leave:
			// 退室
			delete(r.Clients, client)
			close(client.Send)
			r.Tracer.Trace("クライアントが退室しました")
		case msg := <-r.Forward:
			r.Tracer.Trace("メッセージを受信しました: ", msg.Message)
			// すべてのクライアントにメッセージを転送
			for client := range r.Clients {
				select {
				case client.Send <- msg:
					// メッセージを送信
					r.Tracer.Trace(" -- クライアントに送信されました")
				default:
					// 送信に失敗
					delete(r.Clients, client)
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

	// create Client
	userData, err := c.Cookie("username")
	if err != nil {
		return err
	}
	client := &Client{
		ID:       1,
		Name:     "test",
		Socket:   ws,
		Send:     make(chan *Message, messageBufferSize),
		Room:     room,
		UserData: userData,
	}
	fmt.Printf("%#v\n", rooms)

	// client Join Room
	room.Join <- client
	defer func() { room.Leave <- client }()
	go client.Write()
	client.Read()
	return nil
}
