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

	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	for i, ui := range r.users {
		if ui == nil {
			u.Id = i
			r.users[i] = u
			break
		}
	}

	return nil
}

func (r *userRepository) FindById(ctx context.Context, id int) (*models.User, error) {
	for _, u := range r.users {
		if u == nil {
			continue
		}

		if u.Id == id {
			return u, nil
		}
	}

	return nil, ErrNotFound
}

func (r *userRepository) FindByUsername(ctx context.Context, username, password string) (*models.User, error) {
	for _, u := range r.users {
		if u == nil {
			continue
		}

		if u.Username == username && u.ComparePassword(password) {
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

		if u.Email == email && u.ComparePassword(password) {
			return u, nil
		}
	}

	return nil, ErrNotFound
}
