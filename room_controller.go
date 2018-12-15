package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/anraku/chat/domain"
	"github.com/anraku/chat/interfaces"
	"github.com/anraku/chat/trace"
	"github.com/anraku/chat/usecase"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

// package interfaces

type RoomController struct {
	RoomInteractor    *usecase.RoomInteractor
	MessageInteractor *usecase.MessageInteractor
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
	rooms = make(map[string]*domain.Room, 1000)
)

// Index render list of chat room
func (controller *RoomController) Index(c interfaces.Context) error {
	rooms, err := controller.RoomInteractor.Fetch()
	if err != nil {
		return err
	}
	// get username from context
	var username string
	if v, ok := c.Get("username").(string); ok {
		username = v
	} else {
		username = ""
	}
	m := map[string]interface{}{
		"username": username,
		"rooms":    rooms,
	}
	return c.Render(http.StatusOK, "index.html", m)
}

// NewRoom render window to create new chat room
func (controller *RoomController) NewRoom(c echo.Context) error {
	return c.Render(http.StatusOK, "new.html", nil)
}

// Room render chat window
func (controller *RoomController) EnterRoom(c interfaces.Context) error {
	req := c.Request()
	uri := "ws://" + req.Host
	roomID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	messages, err := controller.MessageInteractor.GetByRoomID(roomID)
	if err != nil {
		return err
	}
	data := map[string]interface{}{
		"ID":       roomID,
		"Uri":      uri,
		"Messages": messages,
	}
	return c.Render(http.StatusOK, "chat.html", data)
}

// CreateRoom create new room
func (controller *RoomController) CreateRoom(c interfaces.Context) error {
	newRoom := domain.Room{
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
	}
	err := controller.RoomInteractor.Create(newRoom)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return c.Redirect(http.StatusMovedPermanently, "/index")
}

// Chat is Handler with WebSocket in chat room
func (*RoomController) Chat(c interfaces.Context) error {
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
	var room *domain.Room
	if _, ok := rooms[roomID]; ok {
		room = rooms[roomID]
	} else {
		room = domain.NewRoom(id)
		room.Tracer = trace.New(os.Stdout)
		room.ID = id
		rooms[roomID] = room
		go room.Run()
	}

	// get username from context
	var username string
	if v, ok := c.Get("username").(string); ok {
		username = v
	} else {
		username = ""
	}
	client := &domain.User{
		ID:     1,
		Name:   username,
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
