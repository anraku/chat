package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/anraku/chat/domain/model"
	"github.com/anraku/chat/usecase"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

type RoomController struct {
	RoomInteractor    usecase.RoomUsecase
	MessageInteractor usecase.MessageUsecase
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
	rooms = make(map[string]*model.Room, 1000)
)

// Index render list of chat room
func (controller *RoomController) Index(c echo.Context) error {
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
	data := map[string]interface{}{
		"username": username,
		"rooms":    rooms,
	}
	return c.Render(http.StatusOK, "index.html", data)
}

// NewRoom render window to create new chat room
func (controller *RoomController) NewRoom(c echo.Context) error {
	return c.Render(http.StatusOK, "new.html", nil)
}

// Room render chat window
func (controller *RoomController) EnterRoom(c echo.Context) error {
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
func (controller *RoomController) CreateRoom(c echo.Context) error {
	newRoom := model.Room{
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
func (controller *RoomController) Chat(c echo.Context) error {
	// setting WebSocket
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	// get roomID from URL parameter
	roomID, err := strconv.Atoi(c.Param("id"))
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
	user := &model.User{
		Name:   username,
		Socket: ws,
		Send:   make(chan *model.Message, messageBufferSize),
	}

	room := model.GetRoom(roomID)
	controller.MessageInteractor.EnterRoom(user, room)
	return nil
}
