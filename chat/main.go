package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/anraku/chat/trace"
	"github.com/gobuffalo/packr"
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
		box := packr.NewBox("./templates")
		html, err := box.FindString(t.filename)
		if err != nil {
			panic(err)
		}
		t.templ = template.New(t.filename)
		t.templ, err = t.templ.Parse(html)
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	t.templ.Execute(w, data)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse() // フラグを解釈します

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
	log.Println("Webサーバーを開始します。ポート: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
