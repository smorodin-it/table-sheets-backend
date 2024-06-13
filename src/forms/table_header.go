package forms

type TableHeader struct {
	Label    string `db:"label" json:"label"`
	ParentID string `db:"parent_id" json:"parentId"`
}
