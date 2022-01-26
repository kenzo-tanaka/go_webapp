package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/kenzo-takana/go_webapp/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "application address")
	flag.Parse()

	// TODO: .env管理
	gomniauth.SetSecurityKey("key")
	gomniauth.WithProviders(
		google.New("client id", "secret key", "http://localhost:8080/auth/callback/google"),
	)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	// *authHandlerのServeHTTPが実行され、
	// 認証に成功したら*templateHandlerが実行される
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)

	go r.run()
	log.Println("web server on", *addr)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
