package repositories

import (
	"github.com/jmoiron/sqlx"
	"min-selhoz-backend/src/domains"
)

type TableCell2TableCell interface {
	ListByValueID(tableCellValueID string) (*[]domains.TableCell2TableCell, error)
	Create(tCell2TCell domains.TableCell2TableCell) (*domains.TableCell2TableCell, error)
	Delete(id string) error
}

type TableCell2TableCellRepository struct {
	db *sqlx.DB
}

func (r TableCell2TableCellRepository) ListByValueID(tableCellValueID string) (*[]domains.TableCell2TableCell, error) {
	tCell2TCells := new([]domains.TableCell2TableCell)
	sql := "SELECT * FROM table_cell_2_table_cell WHERE table_cell_value_id=&1"
	err := r.db.Get(&tCell2TCells, sql, tableCellValueID)
	if err != nil {
		return nil, err
	}

	return tCell2TCells, nil
}

func (r TableCell2TableCellRepository) Create(tCell2TCell domains.TableCell2TableCell) (*domains.TableCell2TableCell, error) {
	sql := "INSERT INTO table_cell_2_table_cell (table_cell_value_id, table_cell_argument_id, id, created_at, updated_at) VALUES (:table_cell_value_id, :table_cell_argument_id, :id, :created_at, :updated_at)"
	_, err := r.db.NamedQuery(sql, tCell2TCell)
	if err != nil {
		return nil, err
	}

	return &tCell2TCell, nil
}

func (r TableCell2TableCellRepository) Delete(id string) error {
	sql := "DELETE FROM table_cell_2_table_cell WHERE id=&1"
	_, err := r.db.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}
