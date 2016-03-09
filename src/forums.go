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
	CurrentPage  int
	Topics       []Topic
	PageCount    int
	Pages        []int
}

func (f Forum) GetTopics(page int64) ([]Topic, int) {
	resultCount := 0
	offset := TOPICS_PER_PAGE * (page - 1)
	rows, err := DB.Query("SELECT SQL_CALC_FOUND_ROWS t.topic_id, t.topic_guid, t.topic_title FROM topics t WHERE t.forum_id=? LIMIT ? OFFSET ?", f.ID, TOPICS_PER_PAGE, offset)
	err = DB.QueryRow("SELECT count(*) as count FROM topics t WHERE t.forum_id=?", f.ID).Scan(&resultCount)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("count", resultCount)
	for rows.Next() {
		t := Topic{}
		rows.Scan(&t.ID, &t.Guid, &t.Title)
		f.Topics = append(f.Topics, t)
	}
	return f.Topics, (resultCount / TOPICS_PER_PAGE)
}

func GetForum(guid string, page int64) Forum {
	f := Forum{}
	err := DB.QueryRow("SELECT forum_id, forum_name FROM forums WHERE forum_guid=?", guid).Scan(&f.ID, &f.Name)
	if err != nil {
		fmt.Println(err)
	}
	f.Topics, f.PageCount = f.GetTopics(page)
	for i := 1; i <= f.PageCount; i++ {
		f.Pages = append(f.Pages, i)
	}

	return f
}
