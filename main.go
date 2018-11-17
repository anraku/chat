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

// templは1つのテンプレートを表します
type Template struct {
	// once     sync.Once
	// filename string
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var roomArray = make([]room, 1000, 2000)

func Chat(c echo.Context) error {
	req := c.Request()
	uri := "ws://" + req.Host
	data := map[string]interface{}{
		// "ID":  c.Param("id"),
		"Uri": uri,
	}
	return c.Render(http.StatusOK, "chat.html", data)
}

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

func NewRoom(c echo.Context) error {
	return c.Render(http.StatusOK, "new.html", nil)
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	v := r.Form
	newRoom := domain.Room{
		Name:        v.Get("name"),
		Description: v.Get("description"),
	}
	err = repository.NewRepository(DB).Create(newRoom)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/index", http.StatusFound)
}

var DB *gorm.DB

func main() {
	// Setup db
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	DB = db
	defer db.Close()

	r := newRoom()
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
	e.GET("/chat/:id", Chat)
	e.GET("/room/new", NewRoom)
	http.Handle("/room", r)
	// http.Handle("/room", roomArray[])
	http.HandleFunc("/room/create", CreateRoom)
	// チャットルームを開始します
	go r.run()
	e.Logger.Fatal(e.Start(":8080"))
}
