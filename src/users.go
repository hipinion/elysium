package elysium

import ()

type User struct {
	ID    int    `json:"user_id"`
	Guid  string `json:"user_guid"`
	Name  string `json:"user_name"`
	Email string `json:"user_email"`
}
