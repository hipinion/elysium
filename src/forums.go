package elysium

import (
	"fmt"
)

type Forum struct {
	ID           int    `json:"forum_id"`
	Guid         string `json:"forum_guid"`
	Name         string `json:"forum_name"`
	Organization int    `json:"organization_id"`
	Publication  int    `json:"publication_id"`
	Topics       []Topic
}

func (f Forum) GetTopics() []Topic {
	rows, err := DB.Query("SELECT t.topic_id, t.topic_guid, t.topic_title FROM topics t WHERE t.forum_id=?", f.ID)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		t := Topic{}
		rows.Scan(&t.ID, &t.Guid, &t.Title)
		f.Topics = append(f.Topics, t)
	}
	return f.Topics
}

func GetForum(id string) Forum {
	f := Forum{}
	err := DB.QueryRow("SELECT forum_id, forum_name FROM forums WHERE forum_id=?", id).Scan(&f.ID, &f.Name)
	if err != nil {
		fmt.Println(err)
	}
	f.Topics = f.GetTopics()

	return f
}
