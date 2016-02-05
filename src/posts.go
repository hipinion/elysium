package elysium

import (
	"log"
)

type Post struct {
	Guid string `json:"post_guid"`
	Text string `json:"post_text"`
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
