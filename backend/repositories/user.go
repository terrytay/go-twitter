package repositories

import (
	"context"

	"github.com/terrytay/twitter/backend/entities"
	"github.com/terrytay/twitter/backend/infrastructures/db"
)

type IUserRepository interface {
	Create(ctx context.Context, newUser *entities.User) error
}

type UserRepository struct {
	Db *db.Database
}

func (r *UserRepository) Create(ctx context.Context, newUser *entities.User) error {
	err := r.Db.Client.Create(ctx, "users", newUser)
	if err != nil {
		return err
	}
	return nil
}
