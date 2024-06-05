package repositories

import (
	"github.com/jmoiron/sqlx"
	"min-selhoz-backend/src/domains"
)

type User interface {
	List() (*[]domains.User, error)
	Retrieve(id string) (*domains.User, error)
	Create(*domains.User) (*domains.User, error)
	Update(*domains.User) (*domains.User, error)
	Enable(updateBool *domains.UpdateBool) (*domains.UpdateBool, error)
}

type UserRepository struct {
	db *sqlx.DB
}

func (r UserRepository) List() (*[]domains.User, error) {
	users := new([]domains.User)
	sql := "SELECT * FROM \"user\""
	err := r.db.Select(users, sql)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r UserRepository) Retrieve(id string) (*domains.User, error) {
	user := new(domains.User)
	sql := "SELECT * FROM \"user\" WHERE user_id=$1"
	err := r.db.Get(user, sql, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepository) Create(user *domains.User) (*domains.User, error) {
	sql := "INSERT INTO \"user\" (user_id, username, first_name, last_name, patronymic, enabled, role, organization_id) VALUES (:user_id, :username, :first_name, :last_name, :patronymic, :enabled, :role, :organization_id)"
	_, err := r.db.NamedExec(sql, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepository) Update(user *domains.User) (*domains.User, error) {
	sql := "UPDATE \"user\" SET username=:username, first_name=:first_name, last_name=:last_name, patronymic=:patronymic, role=:role, organization_id=:organization_id WHERE user_id=:user_id"
	_, err := r.db.NamedExec(sql, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepository) Enable(status *domains.UpdateBool) (*domains.UpdateBool, error) {
	sql := "UPDATE \"user\" SET enabled=:status WHERE user_id=:user_id"
	_, err := r.db.NamedExec(sql, status)
	if err != nil {
		return nil, err
	}

	return status, nil
}
