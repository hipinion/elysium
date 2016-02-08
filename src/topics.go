package elysium

import (
	"log"
)

type Topic struct {
	ID    int    `json:"topic_id"`
	Title string `json:"topic_title"`
	Guid  string `json:"topic_guid"`

	User  int    `json:"user_id"`
	Time  int    `json:"topic_create_time"`
	Posts []Post `json:"posts"`
}

func GetTopic(guid string) Topic {
	t := Topic{}
	err := DB.QueryRow("SELECT topic_guid, topic_title FROM topics WHERE topic_guid=?", guid).Scan(&t.Guid, &t.Title)
	if err != nil {
		log.Println(err)
	}
	return t
}
