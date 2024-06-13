package forms

type User struct {
	Username       string `db:"username" json:"username"`
	FirstName      string `db:"first_name" json:"firstName"`
	LastName       string `db:"last_name" json:"lastName"`
	Patronymic     string `db:"patronymic" json:"patronymic"`
	Role           int    `db:"role" json:"role"`
	OrganizationId string `db:"organization_id" json:"organizationId"`
}

func (u User) Validate() error {
	return nil
}
