package usecases

import (
	"errors"

	"github.com/blackriper/manager/domain"
)

type CreateAdmin struct {
	UserRepo domain.RepositoryAdmin
}

func NewCreateAdmin(userRepo domain.RepositoryAdmin) *CreateAdmin {
	return &CreateAdmin{
		UserRepo: userRepo,
	}
}

func (a *CreateAdmin) CreateAdmin(user domain.User) error {
	if a.UserRepo.ExistsAdmin(user.Email) {
		return errors.New("admin already exists")
	}
	err := a.UserRepo.CreateAdmin(user)
	return err
}
