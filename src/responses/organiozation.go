package responses

import "time"

type OrganizationResp struct {
	ID        string    `json:"id"`
	Label     string    `json:"label"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type OrganizationListResp = []OrganizationResp
