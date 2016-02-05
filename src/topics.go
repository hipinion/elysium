package elysium

import (
	"log"
)

type Topic struct {
	ID    int    `json:"topic_id"`
	Guid  string `json:"topic_guid"`
	Title string `json:"topic_title"`
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
