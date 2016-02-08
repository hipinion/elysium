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

func (t Topic) GetPosts() []Post {

	rows, err := DB.Query("SELECT post_id, post_text, user_id FROM posts WHERE topic_id=?", t.ID)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		p := Post{}
		rows.Scan(&p.ID, &p.Text, &p.User)
		t.Posts = append(t.Posts, p)
	}
	return t.Posts
}

func GetTopic(guid string) Topic {
	t := Topic{}
	err := DB.QueryRow("SELECT topic_id, topic_guid, topic_title FROM topics WHERE topic_guid=?", guid).Scan(&t.ID, &t.Guid, &t.Title)
	if err != nil {
		log.Println(err)
	}
	t.Posts = t.GetPosts()

	return t
}
