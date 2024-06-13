package services

import (
	"github.com/google/uuid"
	"min-selhoz-backend/src/domains"
	"min-selhoz-backend/src/forms"
	"min-selhoz-backend/src/repositories"
	"min-selhoz-backend/src/responses"
	"time"
)

type UserServiceInterface interface {
	List() (*responses.UserListResp, error)
	Retrieve(id string) (*responses.UserResp, error)
	Create(form *forms.UserForm) (*responses.UserResp, error)
	Update(form *forms.UserForm) (*responses.UserResp, error)
	Enable(form *forms.UpdateBoolForm) error
}

type UserService struct {
	r *repositories.UserRepository
}

func (s UserService) List() (*responses.UserListResp, error) {
	usersDomains, err := s.r.List()
	if err != nil {
		return nil, err
	}

	users := make(responses.UserListResp, 0)
	for _, user := range *usersDomains {
		users = append(users, responses.UserResp{
			ID:             user.ID,
			Username:       user.Username,
			FirstName:      user.FirstName,
			LastName:       user.LastName,
			Patronymic:     user.Patronymic,
			Enabled:        user.Enabled,
			LastLogin:      user.LastLogin,
			Role:           user.Role,
			OrganizationId: user.OrganizationId,
		})
	}

	return &users, nil
}

func (s UserService) Retrieve(id string) (*responses.UserResp, error) {
	usersDomain, err := s.r.Retrieve(id)
	if err != nil {
		return nil, err
	}

	user := responses.UserResp{
		ID:             usersDomain.ID,
		Username:       usersDomain.Username,
		FirstName:      usersDomain.FirstName,
		LastName:       usersDomain.LastName,
		Patronymic:     usersDomain.Patronymic,
		Enabled:        usersDomain.Enabled,
		LastLogin:      usersDomain.LastLogin,
		Role:           usersDomain.Role,
		OrganizationId: usersDomain.OrganizationId,
	}

	return &user, nil
}

func (s UserService) Create(form *forms.UserForm) (*responses.UserResp, error) {
	userDomain := domains.User{
		ID:             uuid.New().String(),
		Username:       form.Username,
		FirstName:      form.FirstName,
		LastName:       form.LastName,
		Patronymic:     form.Patronymic,
		Enabled:        true,
		Role:           form.Role,
		OrganizationId: form.OrganizationId,
	}
	user, err := s.r.Create(&userDomain)
	if err != nil {
		return nil, err
	}

	userResp := responses.UserResp{
		ID:             user.ID,
		Username:       user.Username,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Patronymic:     user.Patronymic,
		Enabled:        user.Enabled,
		LastLogin:      user.LastLogin,
		Role:           user.Role,
		OrganizationId: user.OrganizationId,
	}

	return &userResp, nil
}

func (s UserService) Update(form *forms.UserForm) (*responses.UserResp, error) {
	userDomain := domains.User{
		ID:             uuid.New().String(),
		Username:       form.Username,
		FirstName:      form.FirstName,
		LastName:       form.LastName,
		Patronymic:     form.Patronymic,
		Enabled:        true,
		Role:           form.Role,
		OrganizationId: form.OrganizationId,
	}
	user, err := s.r.Update(&userDomain)
	if err != nil {
		return nil, err
	}

	userResp := responses.UserResp{
		ID:             user.ID,
		Username:       user.Username,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Patronymic:     user.Patronymic,
		Enabled:        user.Enabled,
		LastLogin:      user.LastLogin,
		Role:           user.Role,
		OrganizationId: user.OrganizationId,
	}

	return &userResp, nil
}

func (s UserService) Enable(form *forms.UpdateBoolForm) error {
	statusDomain := domains.UpdateBool{
		ID:        form.ID,
		Status:    form.Status,
		UpdatedAt: time.Now(),
	}

	err := s.r.Enable(&statusDomain)
	if err != nil {
		return err
	}

	return nil
}

func NewUserService(r *repositories.UserRepository) UserServiceInterface {
	return UserService{r}
}
