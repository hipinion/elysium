package main

import (
	"encoding/json"
	Elysium "github.com/hipinion/elysium/src"
	"io/ioutil"
	"log"
)

const (
	VERSION = "0.01"
	CONFIG  = "config/config.json"
	DATA    = "tools/sample.data.json"
)

func init() {
	log.Println("\360\237\215\224 \tElysium v" + VERSION)
}

type ImportData struct {
	Organizations []Elysium.Organization `json:"organizations"`
	Publications  []Elysium.Publication  `json:"publications"`
	Forums        []Elysium.Forum        `json:"forums"`
	Topics        []Elysium.Topic        `json:"topics"`
	Posts         []Elysium.Post         `json:"posts"`
	Users         []Elysium.User         `json:"users"`
}

func main() {
	Elysium.Init(CONFIG)
	file, err := ioutil.ReadFile(DATA)
	if err != nil {
		log.Println(err)
	}
	data := new(ImportData)
	json.Unmarshal(file, data)
	for _, o := range data.Organizations {
		Elysium.DB.Exec("INSERT INTO organizations SET organization_id=?, organization_name=?, organization_guid=?", o.ID, o.Name, o.Guid)
	}
	for _, p := range data.Publications {
		Elysium.DB.Exec("INSERT INTO publications SET publication_id=?, publication_name=?, publication_guid=?, organization_id=?", p.ID, p.Name, p.Guid, p.Organization)
	}
	for _, f := range data.Forums {
		Elysium.DB.Exec("INSERT INTO forums SET forum_id=?, forum_guid=?, forum_name=?, organization_id=?", f.ID, f.Guid, f.Name, f.Organization)
	}
	for _, t := range data.Topics {
		Elysium.DB.Exec("INSERT INTO topics SET topic_id=?, forum_id=?, topic_guid=?, topic_title=?, user_id=?, topic_create_time=?", t.ID, 1, t.Guid, t.Title, t.User, t.Time)
	}
	for _, p := range data.Posts {

		Elysium.DB.Exec("INSERT INTO posts SET post_id=?, topic_id=?, post_guid=?, post_title=?, post_text=?, user_id=?, post_create_time=?", p.ID, p.Topic, "", "", p.Text, p.User, p.Time)
	}
	for _, u := range data.Users {

		Elysium.DB.Exec("INSERT INTO users SET user_id=?, user_name=?, user_email=?", u.ID, u.Name, u.Email)
	}
}
