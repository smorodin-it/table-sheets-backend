package responses

import (
	"time"
)

type UserResp struct {
	ID             string     `json:"id"`
	Username       string     `json:"username"`
	FirstName      string     `json:"firstName"`
	LastName       string     `json:"lastName"`
	Patronymic     string     `json:"patronymic"`
	Enabled        bool       `json:"enabled"`
	LastLogin      *time.Time `json:"lastLogin"`
	Role           int        `json:"role"`
	OrganizationId string     `json:"organizationId"`
}

type UserListResp = []UserResp
