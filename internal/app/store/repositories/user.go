package repositories

import (
	"context"

	"github.com/ilfey/go-back/internal/app/store/models"
)

type UserRepository interface {
	Create(ctx context.Context, u *models.User) error
}
