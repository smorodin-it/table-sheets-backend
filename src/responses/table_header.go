package responses

type TableHeader struct {
	ID        string           `json:"id"`
	Label     string           `json:"label"`
	IsDeleted bool             `json:"isDeleted"`
	ParentID  string           `json:"parentId"`
	Child     *TableHeaderList `json:"child"`
}

type TableHeaderList = []TableHeader
