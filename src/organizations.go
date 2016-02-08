package elysium

import ()

type Organization struct {
	ID   int    `json:"organization_id"`
	Guid string `json:"organization_guid"`
	Name string `json:"organization_name"`
}
