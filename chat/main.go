package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/anraku/chat/database"
	"github.com/anraku/chat/domain"
	"github.com/anraku/chat/repository"
	"github.com/anraku/chat/trace"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"google.golang.org/appengine"
)

// templは1つのテンプレートを表します
type Template struct {
	// once     sync.Once
	// filename string
	templates *template.Template
}

// /chatで使ってる
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var roomArray = make([]room, 1000, 2000)

// ServeHTTPはHTTPリクエストを処理します
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
				t.filename)))
	})
	var uri string
	if appengine.IsAppEngine() {
		c := appengine.NewContext(r)
		uri = "wss://" + appengine.DefaultVersionHostname(c)
	} else {
		uri = "ws://" + r.Host
	}
	data := map[string]interface{}{
		"ID":  r.Form.Get("room_id"),
		"Uri": uri,
	}
	if err := t.templ.Execute(w, data); err != nil {
		panic(err)
	}
}

func Index(c echo.Context) error {
	rooms, err := repository.NewRepository(DB).Fetch()
	if err != nil {
		panic(err)
	}

	m := map[string]interface{}{
		"rooms": rooms,
	}
	// t := template.Must(template.ParseFiles("templates/index.html"))
	// if err := t.ExecuteTemplate(w, "index.html", m); err != nil {
	// 	log.Fatal(err)
	// }
	return c.Render(http.StatusOK, "index.html", m)
}

func NewRoom(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/new.html"))
	if err := t.ExecuteTemplate(w, "new.html", nil); err != nil {
		log.Fatal(err)
	}
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
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	DB = db
	defer db.Close()

	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	// static file path
	// http.HandleFunc("/vendor/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, r.URL.Path[1:])
	// })
	// http.HandleFunc("/css/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, r.URL.Path[1:])
	// })
	// http.HandleFunc("/avatars/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, r.URL.Path[1:])
	// })

	e := echo.New()
	// setting static files
	e.Static("/vendor", "vendor")
	e.Static("/css", "css")
	e.Static("/avatars", "avatars")

	// setting tmeplates
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t

	e.GET("/index", Index)
	// http.HandleFunc("/index", Index)
	http.Handle("/chat/:id", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// http.Handle("/room", roomArray[])
	http.HandleFunc("/room/new", NewRoom)
	http.HandleFunc("/room/create", CreateRoom)
	// チャットルームを開始します
	go r.run()
	// Webサーバーを起動します
	// if appengine.IsAppEngine() {
	// 	appengine.Main()
	// } else {
	// 	if err := http.ListenAndServe(":8080", nil); err != nil {
	// 		log.Fatal("ListenAndServe:", err)
	// 	}
	// }
	// echoのサーバ起動
	e.Logger.Fatal(e.Start(":8000"))
}
