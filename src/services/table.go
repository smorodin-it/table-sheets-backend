package services

import (
	"github.com/google/uuid"
	"min-selhoz-backend/src/domains"
	"min-selhoz-backend/src/forms"
	"min-selhoz-backend/src/repositories"
	"min-selhoz-backend/src/responses"
	"time"
)

type TableServiceInterface interface {
	List() (*responses.TableList, error)
	Retrieve(id string) (*responses.Table, error)
	Create(table *forms.Table) (*responses.ResponseAdd, error)
	Update(table *forms.Table, id string) (*responses.ResponseStatus, error)
	SetDelete(status *forms.UpdateBool, id string) (*responses.ResponseStatus, error)
}

type TableService struct {
	r *repositories.TableRepository
}

func (s TableService) List() (*responses.TableList, error) {
	tables, err := s.r.List()
	if err != nil {
		return nil, err
	}

	tablesResp := make(responses.TableList, 0)

	for _, table := range *tables {
		tablesResp = append(tablesResp, responses.Table{
			ID:             table.ID,
			Label:          table.Label,
			IsDeleted:      table.IsDeleted,
			OrganizationId: table.OrganizationId,
		})
	}

	return &tablesResp, nil
}

func (s TableService) Retrieve(id string) (*responses.Table, error) {
	table, err := s.r.Retrieve(id)
	if err != nil {
		return nil, err
	}

	tableResp := responses.Table{
		ID:             table.ID,
		Label:          table.Label,
		IsDeleted:      table.IsDeleted,
		OrganizationId: table.OrganizationId,
	}

	return &tableResp, nil
}

func (s TableService) Create(form *forms.Table) (*responses.ResponseAdd, error) {
	tableDomain := domains.Table{
		ID:             uuid.New().String(),
		Label:          form.Label,
		IsDeleted:      false,
		OrganizationId: form.OrganizationId,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err := s.r.Create(&tableDomain)
	if err != nil {
		return nil, err
	}

	return &responses.ResponseAdd{ID: tableDomain.ID}, nil
}

func (s TableService) Update(form *forms.Table, id string) (*responses.ResponseStatus, error) {
	tableDomain := domains.TableUpdate{
		ID:        id,
		UpdatedAt: time.Now(),
		Table: forms.Table{
			Label:          form.Label,
			OrganizationId: form.OrganizationId,
		},
	}

	err := s.r.Update(&tableDomain)
	if err != nil {
		return nil, err
	}

	return &responses.ResponseStatus{Status: true}, nil
}

func (s TableService) SetDelete(form *forms.UpdateBool, id string) (*responses.ResponseStatus, error) {
	statusDomain := domains.UpdateBool{
		ID:        id,
		Status:    form.Status,
		UpdatedAt: time.Now(),
	}

	err := s.r.SetDelete(&statusDomain)
	if err != nil {
		return nil, err
	}

	return &responses.ResponseStatus{Status: true}, nil
}

func NewTableService(r *repositories.TableRepository) TableServiceInterface {
	return TableService{r}
}
