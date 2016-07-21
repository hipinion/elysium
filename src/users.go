package elysium

import (
	_ "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	_ "io"
	"log"
	"math/rand"
)

const (
	PW_SALT_LENGTH = 16
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ID       int    `json:"user_id"`
	Guid     string `json:"user_guid"`
	Name     string `json:"user_name"`
	Email    string `json:"user_email"`
	Password string `json:"user_password"`
}

func hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	sha := h.Sum(nil)
	encoded := base64.StdEncoding.EncodeToString([]byte(sha))
	return string(encoded)
}

func generateSalt() string {
	b := make([]rune, PW_SALT_LENGTH)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getSalt(un string) (bool, string) {
	var salt string
	err := DB.QueryRow("SELECT user_salt FROM users WHERE user_name=?", un).Scan(&salt)
	if err != nil {
		return false, ""
	} else {
		return true, salt
	}
}

func (u User) authenticate() bool {
	_, salt := getSalt(u.Name)
	pass := hash(u.Password + salt)
	var count int
	err := DB.QueryRow("SELECT count(*) FROM users WHERE user_name=? AND user_password=? AND user_salt=?", u.Name, pass, salt).Scan(&count)
	if err != nil {
		log.Println(err)
	}
	if count == 1 {
		return true
	} else {
		return false
	}
}

func (u User) create() bool {

	salt := generateSalt()
	pass := hash(u.Password + salt)
	log.Println(salt, pass)
	_, err := DB.Exec("INSERT INTO users SET user_name=?, user_email=?, user_salt=?, user_password=?", u.Name, u.Email, salt, pass)
	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
}
