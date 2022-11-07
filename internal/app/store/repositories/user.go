package repositories

import (
	"context"

	"github.com/ilfey/go-back/internal/app/store/models"
)

type UserRepository interface {
	Create(ctx context.Context, u *models.User) error
	FindById(ctx context.Context, id int) (u *models.User, err error)
	FindByUsername(ctx context.Context, username string) (u *models.User, err error)
	FindByUsernameWithPassword(ctx context.Context, username string, password string) (u *models.User, err error)
	FindByEmail(ctx context.Context, username string) (u *models.User, err error)
	FindByEmailWithPassword(ctx context.Context, email string, password string) (u *models.User, err error)
}
