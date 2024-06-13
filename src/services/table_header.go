package services

import (
	"min-selhoz-backend/src/forms"
	"min-selhoz-backend/src/repositories"
	"min-selhoz-backend/src/responses"
)

type TableHeaderInterface interface {
	ListByTableId(tableID string) (*responses.TableHeaderList, error)
	Retrieve(id string) (*responses.TableHeader, error)
	Create(form *forms.TableHeader) error
	Update(form *forms.TableHeader, id string) error
	SetDelete(form *forms.UpdateBool, id string) error
}

type TableHeaderService struct {
	r *repositories.TableHeaderRepository
}

func (s TableHeaderService) ListByTableId(tableID string) (*responses.TableHeaderList, error) {
	//TODO implement me
	panic("implement me")
}

func (s TableHeaderService) Retrieve(id string) (*responses.TableHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (s TableHeaderService) Create(form *forms.TableHeader) error {
	//TODO implement me
	panic("implement me")
}

func (s TableHeaderService) Update(form *forms.TableHeader, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s TableHeaderService) SetDelete(form *forms.UpdateBool, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewTableHeaderService(r *repositories.TableHeaderRepository) TableHeaderInterface {
	return TableHeaderService{r}
}
