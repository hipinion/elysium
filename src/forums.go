package elysium

import ()

type Forum struct {
	ID           int    `json:"forum_id"`
	Guid         string `json:"forum_guid"`
	Name         string `json:"forum_name"`
	Organization int    `json:"organization_id"`
	Publication  int    `json:"publication_id"`
}
