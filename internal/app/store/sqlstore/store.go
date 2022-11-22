package sqlstore

import (
	"github.com/ilfey/go-back/internal/app/store/repositories"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type Store struct {
	db             *pgx.Conn
	logger         *logrus.Logger
	userRepository *userRepository
}

func New(db *pgx.Conn, logger *logrus.Logger) *Store {
	return &Store{
		db:     db,
		logger: logger,
	}
}

func (s *Store) User() repositories.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &userRepository{
		store: s,
	}

	return s.userRepository
}
