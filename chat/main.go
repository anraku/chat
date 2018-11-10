package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/anraku/chat/database"
	"github.com/anraku/chat/repository"
	"github.com/anraku/chat/trace"
	"github.com/jinzhu/gorm"
	"google.golang.org/appengine"
)

// templは1つのテンプレートを表します
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTPはHTTPリクエストを処理します
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
		"Uri": uri,
	}
	if err := t.templ.Execute(w, data); err != nil {
		panic(err)
	}
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
	http.HandleFunc("/vendor/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/css/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/avatars/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.HandleFunc("/index", Index)
	http.Handle("/chat", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// チャットルームを開始します
	go r.run()
	// Webサーバーを起動します
	if appengine.IsAppEngine() {
		appengine.Main()
	} else {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	rooms, err := repository.NewRepository(DB).Fetch()
	if err != nil {
		panic(err)
	}

	m := map[string]interface{}{
		"rooms": rooms,
	}
	t := template.Must(template.ParseFiles("templates/index.html"))
	if err := t.ExecuteTemplate(w, "index.html", m); err != nil {
		log.Fatal(err)
	}
}
