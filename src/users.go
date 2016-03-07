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

func (u User) Create() {

	salt := generateSalt()
	pass := hash(u.Password + salt)
	log.Println(salt, pass)
	_, err := DB.Exec("INSERT INTO users SET user_name=?, user_email=?, user_salt=?, user_password=?", u.Name, u.Email, salt, pass)
	if err != nil {
		log.Println(err)
	}
}
