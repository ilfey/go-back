package teststore

import (
	"github.com/ilfey/go-back/internal/app/store/models"
	"github.com/ilfey/go-back/internal/app/store/repositories"
)

type Store struct {
	userRepository *userRepository
}

func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() repositories.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &userRepository{
		store: s,
		users: make([]*models.User, 5),
	}

	return s.userRepository
}
