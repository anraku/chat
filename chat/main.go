package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/anraku/chat/trace"
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
		// "Host": r.Host,
		"Uri": uri,
	}
	if err := t.templ.Execute(w, data); err != nil {
		panic(err)
	}

}

func main() {
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", &templateHandler{filename: "chat.html"})
	http.Handle("/avatars/",
		http.StripPrefix("/avatars/",
			http.FileServer(http.Dir("./avatars"))))

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
