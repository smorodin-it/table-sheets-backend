package repositories

import (
	"github.com/jmoiron/sqlx"
	"min-selhoz-backend/src/domains"
)

type Table interface {
	List() (*[]domains.Table, error)
	Retrieve(id string) (*domains.Table, error)
	Create(table *domains.Table) (*domains.Table, error)
	Update(table *domains.Table) (*domains.Table, error)
	Delete(status *domains.UpdateBool) (*domains.Table, error)
}
type TableRepository struct {
	db *sqlx.DB
}

func (r TableRepository) List() (*[]domains.Table, error) {
	tables := new([]domains.Table)
	sql := "SELECT * FROM \"table\""
	err := r.db.Select(&tables, sql)
	if err != nil {
		return nil, err
	}

	return tables, nil
}

func (r TableRepository) Retrieve(id string) (*domains.Table, error) {
	table := new(domains.Table)
	sql := "SELECT * FROM \"table\" WHERE table_id=$1"
	err := r.db.Get(table, sql, id)
	if err != nil {
		return nil, err
	}

	return table, nil
}

func (r TableRepository) Create(table *domains.Table) (*domains.Table, error) {
	sql := "INSERT INTO \"table\" (table_id, label, is_deleted, organization_id, created_at, updated_at) VALUES (:table_id, :label, :is_deleted, :organization_id, :created_at, :updated_at)"
	_, err := r.db.NamedExec(sql, table)
	if err != nil {
		return nil, err
	}

	return table, nil
}

func (r TableRepository) Update(table *domains.Table) (*domains.Table, error) {
	sql := "UPDATE \"table\" SET label=:label, updated_at=:updated_at WHERE table_id=:table_id"
	_, err := r.db.NamedQuery(sql, table)
	if err != nil {
		return nil, err
	}

	return table, nil
}

func (r TableRepository) Delete(status *domains.UpdateBool) error {
	sql := "UPDATE \"table\" SET is_deleted=:status, updated_at=:updated_at WHERE table_id=:entity_id"
	_, err := r.db.NamedExec(sql, status)
	if err != nil {
		return err
	}

	return nil
}
