package elysium

import ()

type Publication struct {
	ID           int    `json:"publication_id"`
	Organization int    `json:"organization_id"`
	Name         string `json:"publication_name"`
	Guid         string `json:"publication_guid"`
}
