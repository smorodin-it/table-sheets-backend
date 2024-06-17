package repositories

import (
	"github.com/jmoiron/sqlx"
	"min-selhoz-backend/src/domains"
)

type TableHeaderInterface interface {
	ListByTableId(tableID string) (*[]domains.TableHeader, error)
	Retrieve(id string) (*domains.TableHeader, error)
	Create(tHeader *domains.TableHeader) error
	Update(tHeader *domains.TableHeader) error
	SetDelete(status *domains.UpdateBool) error
}

type TableHeaderRepository struct {
	db *sqlx.DB
}

func (r TableHeaderRepository) ListByTableId(tableID string) (*[]domains.TableHeader, error) {
	tHeaders := new([]domains.TableHeader)
	sql := `SELECT *, (COUNT(parent.label) - 1) as level 
            FROM table_header as node, table_header as parent 
            WHERE node.table_id=&1 AND parent.table_id=&1 
            AND node.lft BETWEEN parent.lft AND parent.rgt
            GROUP BY node.label 
            ORDER BY node.lft`

	err := r.db.Select(&tHeaders, sql, tableID)
	if err != nil {
		return nil, err
	}

	return tHeaders, nil
}

func (r TableHeaderRepository) Retrieve(id string) (*domains.TableHeader, error) {
	tHeader := new(domains.TableHeader)
	sql := "SELECT * from table_header WHERE table_id=:table_header_id"
	err := r.db.Get(&tHeader, sql, id)
	if err != nil {
		return nil, err
	}

	return tHeader, nil
}

func (r TableHeaderRepository) Create(tHeader *domains.TableHeader) error {
	sql := `BEGIN;

            DROP TABLE IF EXISTS tmp_table;
            CREATE TEMPORARY TABLE tmp_table AS SELECT rgt FROM table_header WHERE parent_id=:parent_id;

			UPDATE table_header SET rgt = rgt + 2 WHERE rgt > (SELECT rgt FROM tmp_table).rgt;
            UPDATE table_header SET lft = lft + 2 WHERE lft > (SELECT rgt FROM tmp_table).rgt;

            INSERT INTO table_header (table_header_id, label, is_deleted, table_id, parent_id, lft, rgt, created_at, updated_at)
            VALUES (:table_header_id, :label, :is_deleted, :table_id, :parent_id, (SELECT rgt FROM tmp_table).rgt + 1, (SELECT rgt FROM tmp_table).rgt + 2, :created_at, :updated_at);
            COMMIT;`

	_, err := r.db.NamedQuery(sql, tHeader)
	if err != nil {
		return err
	}

	return nil
}

func (r TableHeaderRepository) Update(tHeader *domains.TableHeader) error {
	sql := "UPDATE table_header SET label=:label, parent_id=:parent_id, updated_at=:updated_at WHERE table_id=:table_header_id"
	_, err := r.db.NamedQuery(sql, tHeader)
	if err != nil {
		return err
	}

	return nil
}

func (r TableHeaderRepository) SetDelete(status *domains.UpdateBool) error {
	sql := "UPDATE table_header SET is_deleted=:status WHERE table_id=:entity_id"
	_, err := r.db.NamedQuery(sql, status)
	if err != nil {
		return err
	}

	return nil
}

func NewTableHeaderRepository(db *sqlx.DB) TableHeaderInterface {
	return &TableHeaderRepository{db}
}
