package services

import (
	"context"

	"github.com/terrytay/twitter/backend/entities"
	"github.com/terrytay/twitter/backend/repositories"
	"github.com/terrytay/twitter/backend/tools"
)

type IUserService interface {
	Create(ctx context.Context, user entities.NewUser) error
}

type UserService struct {
	Repo repositories.IUserRepository
}

func (u *UserService) Create(ctx context.Context, user entities.NewUser) error {
	hashedPassword, err := tools.HashPassword(user.Password)
	id := tools.GenerateId()
	if err != nil {
		return err
	}

	newUser := &entities.User{Id: id, Username: user.Username, Name: user.Name, HashedPassword: hashedPassword}

	err = u.Repo.Create(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}
