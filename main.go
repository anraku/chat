package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/anraku/chat/database"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

// DB is database connection
var DB *gorm.DB

func LoginMenu(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

func Login(c echo.Context) error {
	// create session data
	userName := c.FormValue("name")
	if userName == "" {
		userName = "名無しさん"
	}
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Values["username"] = userName
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, "/index")
}

func Logout(c echo.Context) error {
	// delete session
	sess, err := session.Get("session", c)
	if err != nil {
		panic(err)
	}
	delete(sess.Values, "username")
	sess.Save(c.Request(), c.Response())
	return c.Render(http.StatusOK, "logout.html", nil)
}

// Room render chat window
func EnterRoom(c echo.Context) error {
	req := c.Request()
	uri := "ws://" + req.Host
	roomID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	messages, err := NewMessageRepository(DB).GetByRoomID(roomID)
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

// Index render list of chat room
func Index(c echo.Context) error {
	rooms, err := NewRoomRepository(DB).Fetch()
	if err != nil {
		panic(err)
	}
	// get username from cookie
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	username := sess.Values["username"]
	// username, _ := url.QueryUnescape(userData.Value)
	m := map[string]interface{}{
		"username": username,
		"rooms":    rooms,
	}
	return c.Render(http.StatusOK, "index.html", m)
}

// NewRoom render window to create new chat room
func NewRoom(c echo.Context) error {
	return c.Render(http.StatusOK, "new.html", nil)
}

// CreateRoom create new room
func CreateRoom(c echo.Context) error {
	newRoom := Room{
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
	}
	err := NewRoomRepository(DB).Create(newRoom)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return c.Redirect(http.StatusMovedPermanently, "/index")
}

func CheckLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get username from session
		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}
		if sess.Values["username"] == nil || sess.Values["username"] == "" {
			return c.Redirect(http.StatusMovedPermanently, "/login")
		}
		return next(c)
	}
}

func main() {
	// Setup db
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	DB = db
	defer db.Close()

	app := NewRouter()
	app.Start(":8080")
}
