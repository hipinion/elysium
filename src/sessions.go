package elysium

import (
	"github.com/gorilla/securecookie"

	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

type Session struct {
	ID string
}

var hashKey = []byte("very-secret")
var blockKey = []byte("a-lot-secret")
var s = securecookie.New(hashKey, blockKey)

func createSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func CreateSession() string {
	return createSessionID()
}

func GetSession(r *http.Request) User {
	u := User{}
	s, _ := r.Cookie("elysium_sid")
	if s.Value != "" {
		err := DB.QueryRow("SELECT u.user_id,  u.user_name, u.user_email FROM sessions s LEFT JOIN users u ON u.user_id=s.user_id WHERE s.session_id=?", s.Value).Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			fmt.Println(err)
		}
		if u.ID != 0 {
			u.LoggedIn = true
		}

	}
	return u
}

func SaveSession(sess string, u User) {
	_, err := DB.Exec("INSERT INTO sessions SET session_id=?, user_id=?", sess, u.ID)
	if err != nil {

	} else {

	}
}
