package domains

import "time"

type User struct {
	ID             string    `db:"user_id"`
	Username       string    `db:"username"`
	FirstName      string    `db:"first_name"`
	LastName       string    `db:"last_name"`
	Patronymic     string    `db:"patronymic"`
	Enabled        bool      `db:"enabled"`
	LastLogin      time.Time `db:"last_login"`
	Role           int       `db:"role"`
	OrganizationId string    `db:"organization_id"`
}
