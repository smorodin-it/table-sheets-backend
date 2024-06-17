package domains

import (
	"min-selhoz-backend/src/forms"
	"time"
)

type Table struct {
	ID             string    `db:"table_id"`
	Label          string    `db:"label"`
	IsDeleted      bool      `db:"is_deleted"`
	OrganizationId string    `db:"organization_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type TableUpdate struct {
	ID        string    `db:"table_id"`
	UpdatedAt time.Time `db:"updated_at"`
	forms.Table
}

type TableHeader struct {
	ID        string    `db:"table_header_id"`
	Label     string    `db:"label"`
	IsDeleted bool      `db:"is_deleted"`
	TableID   string    `db:"table_id"`
	ParentID  string    `db:"parent_id"`
	Lft       int       `db:"lft"`
	Rgt       int       `db:"rgt"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type TableHeaderUpdate struct {
	ID        string    `db:"table_header_id"`
	UpdatedAt time.Time `db:"updated_at"`
	forms.TableHeader
}

type TableRow struct {
	ID        string    `db:"table_row_id"`
	Label     *string   `db:"label"`
	IsDeleted bool      `db:"is_deleted"`
	TableID   string    `db:"table_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type TableCell struct {
	ID            string    `db:"table_cell_id"`
	TableHeaderID string    `db:"table_header_id"`
	TableRowID    string    `db:"table_row_id"`
	TableID       string    `db:"table_id"`
	Value         *string   `db:"value"`
	Type          int       `db:"type"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type TableCell2TableCell struct {
	TableCellValueID    string    `db:"table_cell_value_id"`
	TableCellArgumentID string    `db:"table_cell_argument_id"`
	ID                  string    `db:"id"`
	CreatedAt           time.Time `db:"created_at"`
	UpdatedAt           time.Time `db:"updated_at"`
}
