package repositories

import (
	"github.com/jmoiron/sqlx"
	"min-selhoz-backend/src/domains"
)

type Organization interface {
	List() (*[]domains.Organization, error)
	Retrieve(id string) (*domains.Organization, error)
	Create(*domains.Organization) (*domains.Organization, error)
	Update(*domains.Organization) (*domains.Organization, error)
	Delete(status *domains.UpdateBool) error
}

type OrganizationRepository struct {
	db *sqlx.DB
}

func (r OrganizationRepository) List() (*[]domains.Organization, error) {
	organizations := new([]domains.Organization)
	sql := "SELECT * FROM organization"
	err := r.db.Select(organizations, sql)
	if err != nil {
		return nil, err
	}

	return organizations, nil
}

func (r OrganizationRepository) Retrieve(id string) (*domains.Organization, error) {
	organization := new(domains.Organization)
	sql := "SELECT * FROM organization WHERE organization_id=$1"
	err := r.db.Get(organization, sql, id)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (r OrganizationRepository) Create(organization *domains.Organization) (*domains.Organization, error) {
	sql := "INSERT INTO organization (organization_id, label, created_at, updated_at) VALUES (:organization_id, :label, :created_at, :updated_at)"
	_, err := r.db.NamedQuery(sql, organization)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (r OrganizationRepository) Update(organization *domains.Organization) (*domains.Organization, error) {
	_, err := r.db.NamedQuery("UPDATE organization SET label=:label, updated_at=:updated_at WHERE organization_id=:organization_id", organization)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (r OrganizationRepository) Delete(status *domains.UpdateBool) error {
	_, err := r.db.NamedQuery("UPDATE organization SET is_deleted=:status, updated_at=:updated_at WHERE organization_id=:organization_id", status)
	if err != nil {
		return err
	}

	return nil
}
