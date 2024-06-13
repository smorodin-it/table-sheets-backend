package services

import (
	"github.com/google/uuid"
	"min-selhoz-backend/src/domains"
	"min-selhoz-backend/src/forms"
	"min-selhoz-backend/src/repositories"
	"min-selhoz-backend/src/responses"
	"time"
)

type OrganizationServiceInterface interface {
	List() (*responses.OrganizationListResp, error)
	Retrieve(id string) (*responses.OrganizationResp, error)
	Create(organization *forms.Organization) (*responses.ResponseAdd, error)
	Update(organization *forms.Organization, id string) (*responses.ResponseStatus, error)
	Delete(status *forms.UpdateBoolForm, id string) (*responses.ResponseStatus, error)
}

type OrganizationService struct {
	r *repositories.OrganizationRepository
}

func (s OrganizationService) List() (*responses.OrganizationListResp, error) {
	organizations, err := s.r.List()
	if err != nil {
		return nil, err
	}

	organizationsResp := make(responses.OrganizationListResp, 0)
	for _, organization := range *organizations {
		organizationsResp = append(organizationsResp, responses.OrganizationResp{
			ID:        organization.ID,
			Label:     organization.Label,
			Enabled:   organization.Enabled,
			CreatedAt: organization.CreatedAt,
			UpdatedAt: organization.UpdatedAt,
		})
	}

	return &organizationsResp, err
}

func (s OrganizationService) Retrieve(id string) (*responses.OrganizationResp, error) {
	organization, err := s.r.Retrieve(id)
	if err != nil {
		return nil, err
	}

	organizationResp := responses.OrganizationResp{
		ID:        organization.ID,
		Label:     organization.Label,
		Enabled:   organization.Enabled,
		CreatedAt: organization.CreatedAt,
		UpdatedAt: organization.UpdatedAt,
	}

	return &organizationResp, err
}

func (s OrganizationService) Create(form *forms.Organization) (*responses.ResponseAdd, error) {
	organizationDomain := domains.Organization{
		ID:        uuid.New().String(),
		Label:     form.Label,
		Enabled:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.r.Create(&organizationDomain)
	if err != nil {
		return nil, err
	}

	return &responses.ResponseAdd{ID: organizationDomain.ID}, err
}

func (s OrganizationService) Update(form *forms.Organization, id string) (*responses.ResponseStatus, error) {
	organizationDomain := domains.OrganizationUpdate{
		ID:        id,
		UpdatedAt: time.Now(),
		Organization: forms.Organization{
			Label: form.Label,
		},
	}
	err := s.r.Update(&organizationDomain)
	if err != nil {
		return nil, err
	}

	return &responses.ResponseStatus{Status: true}, nil
}

func (s OrganizationService) Delete(form *forms.UpdateBoolForm, id string) (*responses.ResponseStatus, error) {
	statusDomain := domains.UpdateBool{
		ID:        id,
		Status:    form.Status,
		UpdatedAt: time.Now(),
	}
	err := s.r.SetEnabled(&statusDomain)
	if err != nil {
		return nil, err
	}

	return &responses.ResponseStatus{Status: true}, nil
}
