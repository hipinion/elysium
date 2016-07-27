package elysium

import (
	"github.com/gorilla/securecookie"

	"crypto/rand"
	"encoding/base64"
	"io"
	_ "net/http"
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

func SaveSession(sess string, u User) {
	_, err := DB.Exec("INSERT INTO sessions SET session_id=?, user_id=?", sess, u.ID)
	if err != nil {

	} else {

	}
}

func GetSession(sess string) User {
	u := User{}
	return u
}
