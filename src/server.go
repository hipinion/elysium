package elysium

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "gapearth")
}
func ForumHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "gapearth")
}
func TopicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "gapearth")
}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "gapearth")
}
func UserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "gapearth")
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "gapearth")
}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "gapearth")
}
func Serve() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/register", RegisterHandler)
	r.HandleFunc("/forum/{forum:[0-9a-z-]+}", ForumHandler)
	r.HandleFunc("/topic/{topic:[0-9a-z-]+}", TopicHandler)
	r.HandleFunc("/post/{post:[0-9a-z-]+}", PostHandler)
	r.HandleFunc("/user/{user:[0-9a-z-]+}", UserHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8083", nil)
}

func Ping() {
	log.Println("PING")
}