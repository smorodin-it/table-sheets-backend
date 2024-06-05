package domains

import "time"

type Organization struct {
	ID        string    `db:"organization_id"`
	Label     string    `db:"label"`
	isDeleted bool      `db:"is_deleted"`
	createdAt time.Time `db:"created_at"`
}
