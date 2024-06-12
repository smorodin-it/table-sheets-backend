package repositories

import (
	"github.com/jmoiron/sqlx"
	"min-selhoz-backend/src/domains"
)

type TableHeader interface {
	List() (*[]domains.TableHeader, error)
	Retrieve(id string) (*domains.TableHeader, error)
	Create(tHeader *domains.TableHeader) (*domains.TableHeader, error)
	Update(tHeader *domains.TableHeader) (*domains.TableHeader, error)
	Delete(status *domains.UpdateBool) error
}

type TableHeaderRepository struct {
	db *sqlx.DB
}

func (r TableHeaderRepository) List() (*[]domains.TableHeader, error) {
	tHeaders := new([]domains.TableHeader)
	sql := "SELECT * from table_header"
	err := r.db.Select(&tHeaders, sql)
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

func (r TableHeaderRepository) Create(tHeader *domains.TableHeader) (*domains.TableHeader, error) {
	sql := "INSERT INTO table_header (table_header_id, label, is_deleted, table_id, parent_id, created_at, updated_at) VALUES (:table_header_id, :label, :is_deleted, :table_id, :parent_id, :created_at, :updated_at)"
	_, err := r.db.NamedExec(sql, tHeader)
	if err != nil {
		return nil, err
	}

	return tHeader, nil
}

func (r TableHeaderRepository) Update(tHeader *domains.TableHeader) (*domains.TableHeader, error) {
	sql := "UPDATE table_header SET label=:label, parent_id=:parent_id, updated_at=:updated_at WHERE table_id=:table_header_id"
	_, err := r.db.NamedExec(sql, tHeader)
	if err != nil {
		return nil, err
	}

	return tHeader, nil
}

func (r TableHeaderRepository) Delete(status *domains.UpdateBool) error {
	sql := "UPDATE table_header SET is_deleted=:status WHERE table_id=:entity_id"
	_, err := r.db.NamedExec(sql, status)
	if err != nil {
		return err
	}

	return nil
}