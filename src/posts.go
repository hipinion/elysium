package elysium

import (
	"html/template"
	"log"
)

type Post struct {
	ID    int           `json:"post_id"`
	User  int           `json:"user_id"`
	Topic int           `json:"topic_id"`
	Guid  string        `json:"post_guid"`
	Text  string        `json:"post_text"`
	FText template.HTML `json:"post_ftext"`
	Time  int           `json:"post_create_time"`
}

type Posts []Post

func GetPost(guid string) Post {
	p := Post{}
	err := DB.QueryRow("SELECT post_guid, post_text FROM posts WHERE post_guid=?", guid).Scan(&p.Guid, &p.Text)
	if err != nil {
		log.Println(err)
	}
	return p
}

func Whoa() {
	log.Println("")
}
