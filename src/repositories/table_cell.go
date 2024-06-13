package repositories

import (
	"github.com/jmoiron/sqlx"
	"min-selhoz-backend/src/domains"
)

type TableCell interface {
	List() (*[]domains.TableCell, error)
	Retrieve(id string) (*domains.TableCell, error)
	Create(tCell domains.TableCell) (*domains.TableCell, error)
	Update(tCell domains.TableCell) (*domains.TableCell, error)
	Delete(status domains.UpdateBool) error
}

type TableCellRepository struct {
	db *sqlx.DB
}

func (r TableCellRepository) List() (*[]domains.TableCell, error) {
	//TODO implement me
	panic("implement me")
}

func (r TableCellRepository) Retrieve(id string) (*domains.TableCell, error) {
	tCell := new(domains.TableCell)
	sql := "SELECT * FROM table_cell WHERE table_id=$1"
	err := r.db.Get(&tCell, sql, id)
	if err != nil {
		return nil, err
	}

	return tCell, nil
}

func (r TableCellRepository) Create(tCell domains.TableCell) (*domains.TableCell, error) {
	sql := "INSERT INTO table_cell (table_cell_id, table_header_id, table_row_id, table_id, value, type, created_at, updated_at) VALUES (:table_cell_id, :table_header_id, :table_row_id, :table_id, :value, :type, :created_at, :updated_at)"
	_, err := r.db.NamedQuery(sql, tCell)
	if err != nil {
		return nil, err
	}

	return &tCell, nil
}

func (r TableCellRepository) Update(tCell domains.TableCell) (*domains.TableCell, error) {
	//TODO implement me
	panic("implement me")
}

func (r TableCellRepository) Delete(status domains.UpdateBool) error {
	//TODO implement me
	panic("implement me")
}
