package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServerHTTP handles the HTTP request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	data := map[string]interface{}{
		"Host": r.Host,
	}

	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	t.templ.Execute(w, data)
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the app.")
	flag.Parse()

	securityKey := os.Getenv("SECURITY_KEY")
	googleID := os.Getenv("GOOGLE_CLIENT_ID")
	googleSecret := os.Getenv("GOOGLE_SECRET_KEY")
	gitID := os.Getenv("GITHUB_CLIENT_ID")
	gitSecret := os.Getenv("GITHUB_SECRET_KEY")

	gomniauth.SetSecurityKey(securityKey)
	gomniauth.WithProviders(
		google.New(googleID, googleSecret, "http://localhost:8080/auth/callback/google"),
		github.New(gitID, gitSecret, "http://localhost:8080/auth/callback/github"),
	)

	r := newRoom()

	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.Handle("/chat", mustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
	http.Handle("/room", r)

	go r.run()

	fmt.Println("😎 Server listenning on: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe :", err)
	}
}
