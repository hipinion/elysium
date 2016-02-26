package elysium

import (
	"fmt"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "gapearth")
}
func ForumHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	t := GetForum(v["forum"])
	Templates.ExecuteTemplate(w, "forum.html", t)
}
func ThreadHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	t := GetTopic(v["topic"])
	Templates.ExecuteTemplate(w, "topic.html", t)
}
func PostHandler(w http.ResponseWriter, r *http.Request) {

}
func UserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "gapearth")
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "gapearth")
}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	Templates.ExecuteTemplate(w, "register.html", map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
	})
}
func Serve() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/register", RegisterHandler)
	r.HandleFunc("/forum/{forum:[0-9a-z-]+}", ForumHandler)
	r.HandleFunc("/topic/{topic:[0-9a-z-]+}", ThreadHandler)
	r.HandleFunc("/post/{post:[0-9a-z-]+}", PostHandler)
	r.HandleFunc("/user/{user:[0-9a-z-]+}", UserHandler)

	// API endpoints
	r.HandleFunc("/api/v1/users", API_v1_UsersHandler)
	r.HandleFunc("/api/v1/users/{user:[0-9a-z-]+}", API_v1_UsersHandler)
	r.HandleFunc("/api/v1/topics", API_v1_TopicsHandler)
	r.HandleFunc("/api/v1/topics/{topic:[0-9a-z-]+}", API_v1_TopicsHandler)
	r.HandleFunc("/api/v1/forums", API_v1_ForumsHandler)
	r.HandleFunc("/api/v1/forums/{forum:[0-9a-z-]+}", API_v1_ForumsHandler)

	r.HandleFunc("/api/v1/posts", API_v1_PostsHandler).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8083", csrf.Protect([]byte("32-byte-long-auth-key"))(r))

}

func Ping() {
	log.Println("PING")
}
