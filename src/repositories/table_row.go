package repositories

import (
	"github.com/jmoiron/sqlx"
	"min-selhoz-backend/src/domains"
)

type TableRow interface {
	List() (*[]domains.TableRow, error)
	Retrieve(id string) (*domains.TableRow, error)
	Create(tRow domains.TableRow) (*domains.TableRow, error)
	Update(tRow domains.TableRow) (*domains.TableRow, error)
	Delete(status domains.UpdateBool) error
}

type TableRowRepository struct {
	db *sqlx.DB
}

func (r TableRowRepository) List() (*[]domains.TableRow, error) {
	tables := new([]domains.TableRow)
	sql := "SELECT * FROM table_row"
	err := r.db.Select(&tables, sql)
	if err != nil {
		return nil, err
	}

	return tables, nil
}

func (r TableRowRepository) Retrieve(id string) (*domains.TableRow, error) {
	table := new(domains.TableRow)
	sql := "SELECT * FROM table_row WHERE table_id=&1"
	err := r.db.Get(&table, sql, id)
	if err != nil {
		return nil, err
	}

	return table, nil
}

func (r TableRowRepository) Create(tRow domains.TableRow) (*domains.TableRow, error) {
	sql := "INSERT INTO table_row (table_row_id, label, is_deleted, table_id, created_at, updated_at) VALUES (:table_row_id, :label, :is_deleted, :table_id, :created_at, :updated_at)"
	_, err := r.db.NamedQuery(sql, tRow)
	if err != nil {
		return nil, err
	}

	return &tRow, nil
}

func (r TableRowRepository) Update(tRow domains.TableRow) (*domains.TableRow, error) {
	sql := "UPDATE table_row SET label=:label WHERE table_row_id=:table_row_id"
	_, err := r.db.NamedExec(sql, tRow)
	if err != nil {
		return nil, err
	}

	return &tRow, nil
}

func (r TableRowRepository) Delete(status domains.UpdateBool) error {
	sql := "UPDATE table_row SET is_deleted=:status, updated_at=:updated_at WHERE table_row_id=:entity_id"
	_, err := r.db.NamedExec(sql, status)
	if err != nil {
		return err
	}

	return nil
}
