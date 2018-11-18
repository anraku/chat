package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/anraku/chat/database"
	"github.com/anraku/chat/domain"
	"github.com/anraku/chat/repository"
	"github.com/anraku/chat/trace"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Template used for creating HTML
type Template struct {
	// once     sync.Once
	// filename string
	templates *template.Template
}

// Render create HTML with template file
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var roomArray = make([]room, 1000, 2000)

// Room render chat window
func Room(c echo.Context) error {
	req := c.Request()
	uri := "ws://" + req.Host
	data := map[string]interface{}{
		// "ID":  c.Param("id"),
		"Uri": uri,
	}
	return c.Render(http.StatusOK, "chat.html", data)
}

// Index render list of chat room
func Index(c echo.Context) error {
	rooms, err := repository.NewRepository(DB).Fetch()
	if err != nil {
		panic(err)
	}

	m := map[string]interface{}{
		"rooms": rooms,
	}
	return c.Render(http.StatusOK, "index.html", m)
}

// NewRoom render window to create new chat room
func NewRoom(c echo.Context) error {
	return c.Render(http.StatusOK, "new.html", nil)
}

// CreateRoom create new room
func CreateRoom(c echo.Context) error {
	newRoom := domain.Room{
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
	}
	err := repository.NewRepository(DB).Create(newRoom)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return c.Redirect(http.StatusMovedPermanently, "/index")
}

// DB is database connection
var DB *gorm.DB
var r *room

func main() {
	// Setup db
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	DB = db
	defer db.Close()

	r = newRoom()
	r.tracer = trace.New(os.Stdout)

	e := echo.New()
	e.Use(middleware.Logger())
	// setting static files
	e.Static("/vendor", "vendor")
	e.Static("/css", "css")
	e.Static("/avatars", "avatars")

	// setting tmeplates
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t

	// Routing
	e.GET("/", Index)
	e.GET("/index", Index)
	e.GET("/chat/:id", Room)
	e.GET("/room", Chat)
	e.GET("/room/new", NewRoom)
	e.POST("/room/create", CreateRoom)
	// http.Handle("/room", roomArray[])
	// チャットルームを開始します
	go r.run()
	e.Logger.Fatal(e.Start(":8080"))
}
