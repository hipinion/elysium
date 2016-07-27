package elysium

import (
	"fmt"
	_ "io"
	"log"
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

func (u *User) authenticate() bool {
	_, salt := getSalt(u.Name)
	pass := hash(u.Password + salt)
	var count int
	err := DB.QueryRow("SELECT count(*), user_id FROM users WHERE user_name=? AND user_password=? AND user_salt=?", u.Name, pass, salt).Scan(&count, &u.ID)
	if err != nil {
		log.Println("uh", err)
	}
	if count == 1 {
		return true
	} else {
		return false
	}
}

func GetUser(guid string) User {
	u := User{}
	fmt.Println(guid)
	err := DB.QueryRow("SELECT u.user_name, u.user_id FROM users u WHERE u.user_name=?", guid).Scan(&u.Name, &u.ID)
	if err != nil {
		log.Println(err)
	}

	return u
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
