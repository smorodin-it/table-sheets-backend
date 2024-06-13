package responses

type Table struct {
	ID             string `json:"id"`
	Label          string `json:"label"`
	IsDeleted      bool   `json:"isDeleted"`
	OrganizationId string `json:"organizationId"`
}

type TableList = []Table
