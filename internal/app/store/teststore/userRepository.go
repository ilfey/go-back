package teststore

import (
	"context"
	"fmt"

	"github.com/ilfey/go-back/internal/app/store/models"
)

type userRepository struct {
	store *Store
	users []*models.User
}

var ErrNotFound = fmt.Errorf("user is not found")

func (r *userRepository) Create(ctx context.Context, u *models.User) error {

	// TODO validate
	for i, ui := range r.users {
		if ui == nil {
			u.Id = i
			r.users[i] = u
			break
		}
	}

	return nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username, password string) (*models.User, error) {
	for _, u := range r.users {
		if u == nil {
			continue
		}
		if u.Username == username && u.Password == password {
			return u, nil
		}
	}

	return nil, ErrNotFound
}

func (r *userRepository) FindByEmail(ctx context.Context, email, password string) (*models.User, error) {
	for _, u := range r.users {
		if u == nil {
			continue
		}
		if u.Email == email && u.Password == password {
			return u, nil
		}
	}

	return nil, ErrNotFound
}
