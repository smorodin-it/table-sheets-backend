package forms

type Table struct {
	Label          string `db:"label" json:"label"`
	OrganizationId string `db:"organization_id" json:"organizationId"`
}
