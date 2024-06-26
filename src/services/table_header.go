package services

import (
	"github.com/google/uuid"
	"min-selhoz-backend/src/domains"
	"min-selhoz-backend/src/forms"
	"min-selhoz-backend/src/repositories"
	"min-selhoz-backend/src/responses"
	"time"
)

type TableHeaderInterface interface {
	ListByTableId(tableID string) (*responses.TableHeaderList, error)
	Retrieve(id string) (*responses.TableHeader, error)
	Create(form *forms.TableHeader) (*string, error)
	Update(form *forms.TableHeader, id string) error
	SetDelete(form *forms.UpdateBool, id string) error
}

type TableHeaderService struct {
	r repositories.TableHeaderInterface
}

func (s TableHeaderService) ListByTableId(tableID string) (*responses.TableHeaderList, error) {
	tableHeaders, err := s.r.ListByTableId(tableID)
	if err != nil {
		return nil, err
	}

	tableHeadersResp := make(responses.TableHeaderList, 0)
	for _, tableHeader := range *tableHeaders {
		tableHeadersResp = append(tableHeadersResp, responses.TableHeader{
			ID:        uuid.New().String(),
			Label:     tableHeader.Label,
			IsDeleted: tableHeader.IsDeleted,
			ParentID:  tableHeader.ParentID,
			//Child:     &tableHeadersResp,
		})
	}

	return &tableHeadersResp, nil
}

func (s TableHeaderService) Retrieve(id string) (*responses.TableHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (s TableHeaderService) Create(form *forms.TableHeader) (*string, error) {
	tableHeader := &domains.TableHeader{
		ID:        uuid.New().String(),
		Label:     form.Label,
		IsDeleted: false,
		TableID:   form.TableID,
		ParentID:  form.ParentID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.r.Create(tableHeader)
	if err != nil {
		return nil, err
	}

	return &tableHeader.ID, nil
}

func (s TableHeaderService) Update(form *forms.TableHeader, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s TableHeaderService) SetDelete(form *forms.UpdateBool, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewTableHeaderService(r repositories.TableHeaderInterface) TableHeaderInterface {
	return &TableHeaderService{r}
}
