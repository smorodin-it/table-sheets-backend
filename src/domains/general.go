package domains

import "time"

type UpdateBool struct {
	Status    bool      `db:"status"`
	UpdatedAt time.Time `db:"updated_at"`
}
