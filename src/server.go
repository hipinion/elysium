package elysium

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	Templates.ExecuteTemplate(w, "index.html", nil)
}

func ForumHandler(w http.ResponseWriter, r *http.Request) {
	User := GetSession(r)
	fmt.Println(User)
	v := mux.Vars(r)
	page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 32)
	if page == 0 {
		page = 1
	}
	t := GetForum(v["forum"], page)
	p := Page{User: User, Payload: t}
	Templates.ExecuteTemplate(w, "forum.html", p)
}

func ThreadHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	t := GetTopic(v["topic"])
	Templates.ExecuteTemplate(w, "topic.html", t)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	d := r.URL.Query()
	us := GetUsers(d)
	Templates.ExecuteTemplate(w, "users.html", us)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	u := GetUser(v["user"])
	Templates.ExecuteTemplate(w, "user.html", u)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	Templates.ExecuteTemplate(w, "login.html", map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
	})
}

func LoginProcess(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userName := r.Form.Get("user_name")
	pass := r.Form.Get("user_pass")
	u := User{Name: userName, Password: pass}
	authenticated := u.authenticate()
	if authenticated {
		fmt.Println(u)
		sess := CreateSession()
		expire := time.Now().AddDate(0, 0, 1)
		cookie := http.Cookie{"elysium_sid", sess, "/", "", expire, expire.Format(time.UnixDate), 86400, false, true, "elysium_sid=" + sess, []string{"elysium_sid=" + sess}}
		http.SetCookie(w, &cookie)
		SaveSession(sess, u)
		Templates.ExecuteTemplate(w, "login_success.html", nil)
	} else {
		Templates.ExecuteTemplate(w, "login_error.html", nil)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	Templates.ExecuteTemplate(w, "register.html", map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
	})
}

func RegisterProcess(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userName := r.Form.Get("user_name")
	email := r.Form.Get("user_email")
	pass := r.Form.Get("user_pass")
	u := User{Name: userName, Email: email, Password: pass}
	created := u.create()
	if !created {
		Templates.ExecuteTemplate(w, "register_error.html", u)
	} else {
		Templates.ExecuteTemplate(w, "register_success.html", u)
	}
}

func Serve() {
	r := mux.NewRouter()

	local := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	r.PathPrefix("/public").Handler(local)

	r.HandleFunc("/", HomeHandler)

	r.HandleFunc("/login", LoginHandler).Methods("GET")

	r.HandleFunc("/login", LoginProcess).Methods("POST")
	r.HandleFunc("/register", RegisterHandler).Methods("GET")
	r.HandleFunc("/register", RegisterProcess).Methods("POST")
	r.HandleFunc("/forum/{forum:[0-9a-z-]+}", ForumHandler)
	r.HandleFunc("/topic/{topic:[0-9a-z-]+}", ThreadHandler)
	r.HandleFunc("/post/{post:[0-9a-z-]+}", PostHandler)
	r.HandleFunc("/user/{user:.+}", UserHandler)
	r.HandleFunc("/users", UsersHandler)
	// r.HandleFunc("/test", TestHandler)

	// API endpoints
	r.HandleFunc("/api/v1/members", API_v1_UsersHandler)
	r.HandleFunc("/api/v1/member/{user:[0-9a-z-]+}", API_v1_UsersHandler)
	r.HandleFunc("/api/v1/topics", API_v1_TopicsHandler)
	r.HandleFunc("/api/v1/topics/{topic:[0-9a-z-]+}", API_v1_TopicsHandler)
	r.HandleFunc("/api/v1/forums", API_v1_ForumsHandler)
	r.HandleFunc("/api/v1/forums/{forum:[0-9a-z-]+}", API_v1_ForumsHandler)
	r.HandleFunc("/api/v1/posts", API_v1_PostsHandler)
	r.HandleFunc("/api/v1/posts/{post:[0-9a-z-]+}", API_v1_PostsHandler)
	r.HandleFunc("/api/v1/posts", API_v1_PostsHandler).Methods("POST")

	http.Handle("/", sessionMiddleware(r))
	http.ListenAndServe(":8083", csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(false))(r))

}

func Ping() {
	log.Println("PING")
}
