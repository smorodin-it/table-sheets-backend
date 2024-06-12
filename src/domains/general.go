package domains

import "time"

type UpdateBool struct {
	ID        string    `db:"entity_id"`
	Status    bool      `db:"status"`
	UpdatedAt time.Time `db:"updated_at"`
}
