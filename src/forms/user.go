package forms

type UserForm struct {
	Username       string `json:"username"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Patronymic     string `json:"patronymic"`
	Role           int    `json:"role"`
	OrganizationId string `json:"organizationId"`
}

func (u UserForm) Validate() error {
	return nil
}
