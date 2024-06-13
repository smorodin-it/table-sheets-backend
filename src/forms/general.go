package forms

type UpdateBoolForm struct {
	ID     string `json:"entityId"`
	Status bool   `json:"status"`
}
