package domains

import (
	"min-selhoz-backend/src/forms"
	"time"
)

type Organization struct {
	ID        string    `db:"organization_id"`
	Label     string    `db:"label"`
	Enabled   bool      `db:"enabled"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type OrganizationUpdate struct {
	ID        string    `db:"organization_id"`
	UpdatedAt time.Time `db:"updated_at"`
	forms.Organization
}
