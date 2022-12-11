package teststore

import (
	"github.com/ilfey/go-back/internal/pkg/store"
	"github.com/ilfey/go-back/internal/pkg/store/models"
)

type Store struct {
	userRepository *userRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &userRepository{
		store: s,
		users: make([]*models.User, 5),
	}

	return s.userRepository
}
